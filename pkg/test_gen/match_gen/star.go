package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenStar(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "StarLiteral",
		Regex:        "a*",
	}, builder.BuildInfo{
		FunctionName: "StarConcatLiteral",
		Regex:        "a*b*",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "star.go"))
}
