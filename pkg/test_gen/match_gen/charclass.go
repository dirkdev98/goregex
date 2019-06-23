package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
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

	formatAndSave(str, err, path.Join(dir, "charclass.go"))
}
