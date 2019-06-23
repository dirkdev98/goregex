package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestRepeaterLiteral(t *testing.T) {
	if ok := match_gen.RepeatLiteral("a"); ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteral("aa"); !ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteral("aaa"); ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteral("bb"); ok {
		t.Fail()
	}
}

func TestRepeaterLiteralMinMax(t *testing.T) {

	if ok := match_gen.RepeatLiteralMinMax("a"); ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteralMinMax("aa"); !ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteralMinMax("aaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteralMinMax("aaaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteralMinMax("aaaaa"); !ok {
		t.Fail()
	}

	if ok := match_gen.RepeatLiteralMinMax("aaaaaa"); ok {
		t.Fail()
	}
}

func TestRepeatWithAlternate(t *testing.T) {
	if ok := match_gen.RepeatWithAlternate("ac"); ok {
		t.Fail()
	}

	if ok := match_gen.RepeatWithAlternate("accc"); ok {
		t.Fail()
	}

	if ok := match_gen.RepeatWithAlternate("ab"); !ok {
		t.Fail()
	}
}

func BenchmarkRepeatLiteral(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.RepeatLiteral("aa")
		benchResult = ok
	}
}

func BenchmarkGoRepeatLiteral(b *testing.B) {
	r, _ := regexp.Compile("a{2}")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("aa")
		benchResult = ok
	}
}
