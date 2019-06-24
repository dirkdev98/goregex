package common_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenDutchZip(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "DutchZip",
		Regex:        "(\\d{4}\\s?[a-zA-Z]{2})+",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "dutch_zip.go"))
}
