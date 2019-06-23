package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
)

func GenAlternate(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "AlternateLiterals",
		Regex:        "ab|cd",
	}, builder.BuildInfo{
		FunctionName: "AlternateMultipleLiterals",
		Regex:        "ab|cd|ef|gh|ij|kl",
	})

	formatAndSave(str, err, path.Join(dir, "alternate.go"))
}
