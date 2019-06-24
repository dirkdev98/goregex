package test_gen

import (
	"fmt"
	"go/format"
	"os"
)

func FormatAndSave(output string, err error, fileName string) {
	if err != nil {
		panic(err)
	}

	bts, err := format.Source([]byte(output))

	if err != nil {
		panic(err)
	}

	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(f, string(bts))
}
