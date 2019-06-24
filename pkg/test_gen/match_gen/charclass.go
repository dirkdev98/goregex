package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenCharClass(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "CharClassSingle",
		Regex:        "[abc]",
	}, builder.BuildInfo{
		FunctionName: "CharClassRange",
		Regex:        "[\\w]",
	}, builder.BuildInfo{
		FunctionName: "CharClassPlusRange",
		Regex:        "[\\d]+",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "charclass.go"))
}
