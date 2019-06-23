package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestLiteralMatch(t *testing.T) {

	if ok := match_gen.LiteralSimple("literal"); !ok {
		t.Fail()
	}

	if ok := match_gen.LiteralSimple("abcdefg"); ok {
		t.Fail()
	}
}

func TestLiteralUnderflow(t *testing.T) {

	if ok := match_gen.LiteralSimple("foo"); ok {
		t.Fail()
	}

	if ok := match_gen.LiteralSimple("litera"); ok {
		t.Fail()
	}
}

func TestLiteralOverflow(t *testing.T) {
	ok := match_gen.LiteralSimple("literalfoo")
	if ok {
		t.Fail()
	}

	if ok := match_gen.LiteralSimple("fooooooooooo"); ok {
		t.Fail()
	}
}

func BenchmarkLiteralExactMatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.LiteralSimple("literal")
		benchResult = ok
	}
}

func BenchmarkGoLiteral(b *testing.B) {
	r, _ := regexp.Compile("literal")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("literal")
		benchResult = ok
	}
}
