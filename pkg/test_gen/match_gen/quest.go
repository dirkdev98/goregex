package match_gen

import (
	"path"

	"github.com/dirkdev98/goregex/pkg/builder"
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

	formatAndSave(str, err, path.Join(dir, "quest.go"))
}
