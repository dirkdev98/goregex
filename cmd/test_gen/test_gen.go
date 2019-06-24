package main

import (
	"os"
	"path"

	"github.com/dirkdev98/goregex/pkg/test_gen/common_gen"
	"github.com/dirkdev98/goregex/pkg/test_gen/match_gen"
)

func main() {
	outdir := "./test/"

	genMatch(outdir)
	genCommon(outdir)
}

func genCommon(outdir string) {
	pkg := "common_gen"
	dir := path.Join(outdir, pkg)
	_ = os.MkdirAll(dir, 0755)

	common_gen.GenDutchZip(pkg, dir)
}

func genMatch(outdir string) {
	pkg := "match_gen"
	dir := path.Join(outdir, pkg)
	_ = os.MkdirAll(dir, 0755)

	match_gen.GenLiterals(pkg, dir)
	match_gen.GenAlternate(pkg, dir)
	match_gen.GenQuest(pkg, dir)
	match_gen.GenRepeat(pkg, dir)
	match_gen.GenPlus(pkg, dir)
	match_gen.GenStar(pkg, dir)
	match_gen.GenCharClass(pkg, dir)
}
