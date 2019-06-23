package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
)

func GenLiterals(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "LiteralSimple",
		Regex:        "literal",
	})

	formatAndSave(str, err, path.Join(dir, "literals.go"))
}
