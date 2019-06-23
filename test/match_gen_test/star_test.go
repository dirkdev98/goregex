package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestStarLiteral(t *testing.T) {
	if ok := match_gen.StarLiteral(""); !ok {
		t.Fail()
	}

	if ok := match_gen.StarLiteral("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarLiteral("aaaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarLiteral("b"); ok {
		t.Fail()
	}
}

func TestStarConcatLiteral(t *testing.T) {
	if ok := match_gen.StarConcatLiteral(""); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("aaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("ab"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("aaaaaab"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("abbbbb"); !ok {
		t.Fail()
	}

	if ok := match_gen.StarConcatLiteral("bb"); !ok {
		t.Fail()
	}
}

func BenchmarkStarConcatLiteral(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.StarConcatLiteral("aaabbb")
		benchResult = ok
	}
}

func BenchmarkGoStarConcatLiteral(b *testing.B) {
	r, _ := regexp.Compile("a*b*")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("aaabbb")
		benchResult = ok
	}
}
