package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
	"github.com/dirkdev98/goregex/pkg/test_gen"
)

func GenQuest(pkg, dir string) {
	str, err := builder.BuildMatchString(pkg, builder.BuildInfo{
		FunctionName: "QuestLiteral",
		Regex:        "ab?",
	}, builder.BuildInfo{
		FunctionName: "QuestDoubleLiteral",
		Regex:        "ab?c?",
	}, builder.BuildInfo{
		FunctionName: "QuestSingleQuest",
		Regex:        "a?",
	})

	test_gen.FormatAndSave(str, err, path.Join(dir, "quest.go"))
}
