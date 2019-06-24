package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenLiterals(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "LiteralSimple",
		Regex:        "literal",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "literals.go"))
}
