package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestPlusLiteral(t *testing.T) {
	if ok := match_gen.PlusLiteral(""); ok {
		t.Fail()
	}

	if ok := match_gen.PlusLiteral("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.PlusLiteral("aaaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.PlusLiteral("b"); ok {
		t.Fail()
	}
}

func TestPlusConcatLiteral(t *testing.T) {
	if ok := match_gen.PlusConcatLiteral(""); ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("a"); ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("aaa"); ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("ab"); !ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("aaaaaab"); !ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("abbbbb"); !ok {
		t.Fail()
	}

	if ok := match_gen.PlusConcatLiteral("bb"); ok {
		t.Fail()
	}
}

func BenchmarkPlusConcatLiteral(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.PlusConcatLiteral("aaabbb")
		benchResult = ok
	}
}

func BenchmarkGoPlusConcatLiteral(b *testing.B) {
	r, _ := regexp.Compile("a+b+")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("aaabbb")
		benchResult = ok
	}
}
