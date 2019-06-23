package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestCharClassSingle(t *testing.T) {
	if ok := match_gen.CharClassSingle("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassSingle("b"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassSingle("c"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassSingle("d"); ok {
		t.Fail()
	}

	if ok := match_gen.CharClassSingle("aa"); ok {
		t.Fail()
	}

	if ok := match_gen.CharClassSingle(""); ok {
		t.Fail()
	}
}

func TestCharClassRange(t *testing.T) {
	if ok := match_gen.CharClassRange("a"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("z"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("A"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("Z"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("_"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("4"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassRange("."); ok {
		t.Fail()
	}
}

func TestCharClassPlusRange(t *testing.T) {
	if ok := match_gen.CharClassPlusRange("1"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassPlusRange("1346"); !ok {
		t.Fail()
	}

	if ok := match_gen.CharClassPlusRange("1aa"); ok {
		t.Fail()
	}

	if ok := match_gen.CharClassPlusRange(""); ok {
		t.Fail()
	}

	if ok := match_gen.CharClassPlusRange("1235807653"); !ok {
		t.Fail()
	}
}

func BenchmarkCharClass(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.CharClassPlusRange("23456789012345678901")
		benchResult = ok
	}
}

func BenchmarkGoCharClass(b *testing.B) {
	r, _ := regexp.Compile("[\\d]+")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("23456789012345678901")
		benchResult = ok
	}
}
