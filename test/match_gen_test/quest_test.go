package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestQuestLiteral(t *testing.T) {
	if ok := match_gen.QuestLiteral("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestLiteral("ab"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestLiteral("abc"); ok {
		t.Fail()
	}
}

func TestQuestDoubleLiteral(t *testing.T) {
	if ok := match_gen.QuestDoubleLiteral("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestDoubleLiteral("ab"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestDoubleLiteral("ac"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestDoubleLiteral("abc"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestDoubleLiteral("abcd"); ok {
		t.Fail()
	}

	if ok := match_gen.QuestDoubleLiteral("abd"); ok {
		t.Fail()
	}
}

func TestQuestSingleQuest(t *testing.T) {
	if ok := match_gen.QuestSingleQuest(""); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestSingleQuest("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.QuestSingleQuest("b"); ok {
		t.Fail()
	}
}

func BenchmarkQuestLiteral(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.QuestLiteral("ab")
		benchResult = ok
	}
}

func BenchmarkGoQuestLiteral(b *testing.B) {
	r, _ := regexp.Compile("ab?")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("ab")
		benchResult = ok
	}
}

func BenchmarkQuestDoubleLiteral(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.QuestDoubleLiteral("abc")
		benchResult = ok
	}
}

func BenchmarkGoQuestDoubleLiteral(b *testing.B) {
	r, _ := regexp.Compile("ab?c?")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("abc")
		benchResult = ok
	}
}
