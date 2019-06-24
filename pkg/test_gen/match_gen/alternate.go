package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenAlternate(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "AlternateLiterals",
		Regex:        "ab|cd",
	}, builder.BuildInfo{
		FunctionName: "AlternateMultipleLiterals",
		Regex:        "ab|cd|ef|gh|ij|kl",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "alternate.go"))
}
