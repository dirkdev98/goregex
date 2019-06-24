package builder

import (
	"fmt"
	"regexp/syntax"
	"strings"
)

type builder struct {
	builder strings.Builder
	pkg     string
}

type funcBuilder struct {
	top     strings.Builder // can be used for function wide variables
	content strings.Builder // anything else
}

type BuildInfo struct {
	FunctionName string
	Regex        string
}

type genInfo struct {
	booleanResultVariable string
	stringIndexVariable   string
}

var uniqueIdx = 0

func getUnique() int {
	result := uniqueIdx
	uniqueIdx++

	return result
}

func BuildMatchString(pkg string, builds ...BuildInfo) (string, error) {
	b := &builder{
		pkg: pkg,
	}
	b.putPackage()

	for _, build := range builds {
		funcBuilder := &funcBuilder{}

		r, err := syntax.Parse(build.Regex, syntax.Perl)

		if err != nil {
			return "", err
		}

		funcBuilder.traverseRegexp(r, genInfo{
			booleanResultVariable: "isMatching",
			stringIndexVariable:   "idx",
		})

		b.prepareFn(&build)

		b.builder.WriteString(funcBuilder.top.String())
		b.builder.WriteString(funcBuilder.content.String())

		b.finalizeFn(&build)
	}

	return b.builder.String(), nil
}

func (b *builder) putPackage() {
	_, _ = fmt.Fprintf(&b.builder, "package %s", b.pkg)
}

func (b *builder) prepareFn(build *BuildInfo) {
	_, _ = fmt.Fprintf(&b.builder, `
// %s
// Generated regex matcher for the following regex
// %s
func %s(src string) bool {
  idx := 0
  isMatching := true

`, build.FunctionName, build.Regex, build.FunctionName)
}

func (b *builder) finalizeFn(build *BuildInfo) {
	_, _ = fmt.Fprintf(&b.builder, `
  if idx != len(src) {
    isMatching = false
  }

 return isMatching
// End of %s
  }
`, build.FunctionName)
}

func (b *funcBuilder) traverseRegexp(reg *syntax.Regexp, info genInfo) {
	switch reg.Op {
	// case syntax.OpEmptyMatch    :
	case syntax.OpLiteral:
		b.addLiteral(reg, info)
	case syntax.OpCharClass:
		b.addCharClass(reg, info)
	// case syntax.OpAnyCharNotNL  :
	// case syntax.OpAnyChar       :
	// case syntax.OpBeginLine     :
	// case syntax.OpEndLine       :
	// case syntax.OpBeginText     :
	// case syntax.OpEndText       :
	// case syntax.OpWordBoundary  :
	// case syntax.OpNoWordBoundary:
	case syntax.OpCapture:
		b.addConcat(reg, info)
	case syntax.OpStar:
		b.addStar(reg, info)
	case syntax.OpPlus:
		b.addPlus(reg, info)
	case syntax.OpQuest:
		b.addQuest(reg, info)
	case syntax.OpRepeat:
		b.addRepeat(reg, info)
	case syntax.OpConcat:
		b.addConcat(reg, info)
	case syntax.OpAlternate:
		b.addAlternate(reg, info)
	default:
		fmt.Println(reg.Op.String(), string(reg.Rune))
		fmt.Printf("%+v\n%+v\n", *reg, info)
	}
}

func (b *funcBuilder) addLiteral(r *syntax.Regexp, info genInfo) {
	uniqueVariable := fmt.Sprintf("staticStr%d", getUnique())

	_, _ = fmt.Fprintf(&b.content, `
  %[1]s := "%s"
  %[1]sLen := len(%[1]s)

  if %[1]sLen + %[4]s > len(src) {
     %[3]s = false
  } else if %[1]s == src[%[4]s:%[4]s + %[1]sLen] {
    %[4]s += %[1]sLen
    %[3]s = true
  } else {
    %[3]s = false
  }
`,
		uniqueVariable,
		string(r.Rune),
		info.booleanResultVariable,
		info.stringIndexVariable,
	)
}

func (b *funcBuilder) addAlternate(r *syntax.Regexp, info genInfo) {
	subLen := len(r.Sub)

	iVar := fmt.Sprintf("i%d", getUnique())
	hasSet := fmt.Sprintf("hasSet%d", getUnique())
	_, _ = fmt.Fprintf(&b.top, "%s := false \n", hasSet)

	_, _ = fmt.Fprintf(&b.content, `
for %s := 1; %s <= %d; %s++ { 
`, iVar, iVar, subLen, iVar)

	for i, t := range r.Sub {
		varName := fmt.Sprintf("alternateResult%d", getUnique())
		_, _ = fmt.Fprintf(&b.top, "%s := false\n", varName)
		_, _ = fmt.Fprintf(&b.content, `
if %s == %d {
`, iVar, i+1)

		b.traverseRegexp(t, genInfo{
			booleanResultVariable: varName,
			stringIndexVariable:   info.stringIndexVariable,
		})

		_, _ = fmt.Fprintf(&b.content, `
	if %s {
		%s = true
		%s = true
		break
	}
}
`, varName, info.booleanResultVariable, hasSet)
	}

	_, _ = fmt.Fprintf(&b.content, `
	}
	if !%s {
		%s = false
	}
`, hasSet, info.booleanResultVariable)
}

func (b *funcBuilder) addConcat(r *syntax.Regexp, info genInfo) {
	// TODO: Create alternative concat that takes captures, stars, +, and repeats in to account

	for i, sub := range r.Sub {
		if i != 0 {
			_, _ = fmt.Fprintf(&b.content, "if %s {\n", info.booleanResultVariable)
		}

		b.traverseRegexp(sub, info)

		if i != 0 {
			b.content.WriteString("\n}\n")
		}
	}
}

func (b *funcBuilder) addQuest(r *syntax.Regexp, info genInfo) {
	boolName := fmt.Sprintf("questBoolResult%d", getUnique())
	idxName := fmt.Sprintf("questIdxResult%d", getUnique())

	b.top.WriteString(fmt.Sprintf("%s := false\n", boolName))
	b.content.WriteString(fmt.Sprintf("%s := %s\n", idxName, info.stringIndexVariable))

	b.traverseRegexp(r.Sub0[0], genInfo{
		booleanResultVariable: boolName,
		stringIndexVariable:   idxName,
	})

	_, _ = fmt.Fprintf(&b.content, `
	if %s {
		%s = %s 
	}
`, boolName, info.stringIndexVariable, idxName)
}

func (b *funcBuilder) addRepeat(r *syntax.Regexp, info genInfo) {
	loopIdxName := fmt.Sprintf("repeatIdx%d", getUnique())
	boolName := fmt.Sprintf("repeatBoolResult%d", getUnique())
	idxName := fmt.Sprintf("repeatIdxResult%d", getUnique())

	subInfo := genInfo{
		booleanResultVariable: boolName,
		stringIndexVariable:   idxName,
	}

	b.top.WriteString(fmt.Sprintf("%s := false\n", boolName))
	b.content.WriteString(fmt.Sprintf("%s := %s\n", idxName, info.stringIndexVariable))

	_, _ = fmt.Fprintf(&b.content, `
for %[1]s := 1; %[1]s <= %[2]d; %[1]s++ {
`, loopIdxName, r.Max, r.Min)

	b.traverseRegexp(r.Sub0[0], subInfo)

	_, _ = fmt.Fprintf(&b.content, `
	if !%[1]s {
		if %[5]s > %[6]d {
			%[4]s = %[2]s
			%[3]s = true
			break
		} else {
			%[3]s = false
			break
		}
	} else if %[1]s && %[5]s == %[7]d {
		%[4]s = %[2]s	
		%[3]s = true
	} 
}
`,
		boolName,
		idxName,
		info.booleanResultVariable,
		info.stringIndexVariable,
		loopIdxName,
		r.Min,
		r.Max,
	)
}

func (b *funcBuilder) addPlus(r *syntax.Regexp, info genInfo) {
	b.traverseRegexp(r.Sub0[0], info)

	boolName := fmt.Sprintf("plusBoolResult%d", getUnique())
	idxName := fmt.Sprintf("plusIdxResult%d", getUnique())

	subInfo := genInfo{
		booleanResultVariable: boolName,
		stringIndexVariable:   idxName,
	}

	b.top.WriteString(fmt.Sprintf("%s := false\n", boolName))
	b.content.WriteString(fmt.Sprintf("%s := %s\n", idxName, info.stringIndexVariable))

	_, _ = fmt.Fprintf(&b.content, `
	if %[1]s {
		for {
`, info.booleanResultVariable)

	b.traverseRegexp(r.Sub0[0], subInfo)

	_, _ = fmt.Fprintf(&b.content, `
			if !%[1]s {
				break
			} else {
				%[3]s = %[1]s
				%[4]s = %[2]s
			}
		}
	}
`, boolName, idxName, info.booleanResultVariable, info.stringIndexVariable)
}

func (b *funcBuilder) addStar(r *syntax.Regexp, info genInfo) {
	boolName := fmt.Sprintf("starBoolResult%d", getUnique())
	idxName := fmt.Sprintf("starIdxResult%d", getUnique())

	subInfo := genInfo{
		booleanResultVariable: boolName,
		stringIndexVariable:   idxName,
	}

	b.top.WriteString(fmt.Sprintf("%s := false\n", boolName))
	b.content.WriteString(fmt.Sprintf("%s := %s\n", idxName, info.stringIndexVariable))

	_, _ = fmt.Fprintf(&b.content, `
		for {
`)

	b.traverseRegexp(r.Sub0[0], subInfo)

	_, _ = fmt.Fprintf(&b.content, `
			if !%[1]s {
				break
			} else {
				%[3]s = %[1]s
				%[4]s = %[2]s
			}
	}
`, boolName, idxName, info.booleanResultVariable, info.stringIndexVariable)
}

func (b *funcBuilder) addCharClass(r *syntax.Regexp, info genInfo) {
	runeVar := fmt.Sprintf("charClassRune%d", getUnique())

	_, _ = fmt.Fprintf(&b.content, `
	if len(src) <= %[2]s {
		%[3]s = false
	} else {
		%[1]s := src[%[2]s]
		if
	`,
		runeVar, info.stringIndexVariable, info.booleanResultVariable)

	runeLength := len(r.Rune)

	for n := 0; n < runeLength; n += 2 {
		if r.Rune[n] == r.Rune[n+1] {
			_, _ = fmt.Fprintf(&b.content, "(%[1]s == %[2]q)", runeVar, r.Rune[n])
		} else {
			_, _ = fmt.Fprintf(&b.content, "(%[1]s >= %[2]q && %[1]s <= %[3]q)", runeVar, r.Rune[n], r.Rune[n+1])
		}

		if n+2 != runeLength {
			b.content.WriteString(" || ")
		}
	}

	_, _ = fmt.Fprintf(&b.content, `{
		%[1]s = %[1]s + 1
		%[2]s = true
	} else {
		%[2]s = false
	}
}
`, info.stringIndexVariable, info.booleanResultVariable)
}
