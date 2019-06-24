package common_gen_test

import (
	"regexp"
	"testing"

	"github.com/dirkdev98/goregex/test/common_gen"
)

func TestDutchZip(t *testing.T) {
	if ok := common_gen.DutchZip(""); ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("1234"); ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("1234AB"); !ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("1234 AB"); !ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("AAAA 12"); ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("1234AB1235AB"); !ok {
		t.Fail()
	}

	if ok := common_gen.DutchZip("1234AC9876 A"); ok {
		t.Fail()
	}
}

func BenchmarkDutchZip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = common_gen.DutchZip("AFTY 23AUEQ12")
	}
}

func BenchmarkGoDutchZip(b *testing.B) {
	r, _ := regexp.Compile("(\\d{4}\\s?[a-zA-Z]{2})+")

	for n := 0; n < b.N; n++ {
		_ = r.MatchString("AFTY 23AUEQ12")
	}
}
