package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenRepeat(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "RepeatLiteral",
		Regex:        "a{2}",
	}, builder.BuildInfo{
		FunctionName: "RepeatLiteralMinMax",
		Regex:        "a{2,5}",
	}, builder.BuildInfo{
		FunctionName: "RepeatWithAlternate",
		Regex:        "ab|ac{2}",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "repeat.go"))
}
