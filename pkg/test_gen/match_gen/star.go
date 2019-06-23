package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
)

func GenStar(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "StarLiteral",
		Regex:        "a*",
	}, builder.BuildInfo{
		FunctionName: "StarConcatLiteral",
		Regex:        "a*b*",
	})

	formatAndSave(str, err, path.Join(dir, "star.go"))
}
