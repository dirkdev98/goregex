package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
)

func GenPlus(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "PlusLiteral",
		Regex:        "a+",
	}, builder.BuildInfo{
		FunctionName: "PlusConcatLiteral",
		Regex:        "a+b+",
	})

	formatAndSave(str, err, path.Join(dir, "plus.go"))
}
