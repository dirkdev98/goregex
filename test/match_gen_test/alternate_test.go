package match_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/match_gen"
)

func TestAlternateMatch(t *testing.T) {
	if ok := match_gen.AlternateLiterals("ab"); !ok {
		t.Fail()
	}

	if ok := match_gen.AlternateLiterals("cd"); !ok {
		t.Fail()
	}

	if ok := match_gen.AlternateLiterals("ef"); ok {
		t.Fail()
	}

	if ok := match_gen.AlternateLiterals("abcd"); ok {
		t.Fail()
	}

	if ok := match_gen.AlternateMultipleLiterals("cd"); !ok {
		t.Fail()
	}

	if ok := match_gen.AlternateMultipleLiterals("gh"); !ok {
		t.Fail()
	}

	if ok := match_gen.AlternateMultipleLiterals("mn"); ok {
		t.Fail()
	}

	if ok := match_gen.AlternateMultipleLiterals("defg"); ok {
		t.Fail()
	}
}

func BenchmarkDoubleAlternate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.AlternateLiterals("ab")
		ok = match_gen.AlternateLiterals("cd")
		benchResult = ok
	}
}

func BenchmarkGoDoubleAlternate(b *testing.B) {
	r, _ := regexp.Compile("ab|cd")

	for n := 0; n < b.N; n++ {
		ok := r.MatchString("ab")
		ok = r.MatchString("cd")
		benchResult = ok
	}
}

func BenchmarkSixthAlternate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ok := match_gen.AlternateMultipleLiterals("kl")
		benchResult = ok
	}
}

func BenchmarkGoSixthAlternate(b *testing.B) {
	r, _ := regexp.Compile("ab|cd|ef|gh|ij|kl")
	for n := 0; n < b.N; n++ {
		ok := r.MatchString("kl")
		benchResult = ok
	}
}
