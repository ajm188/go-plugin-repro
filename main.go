package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/go-plugin-repro/lib"
)

func main() {
	pname := flag.String("plugin", "", "")

	flag.Parse()

	if *pname == "" {
		panic("-plugin cannot be blank")
	}

	fn, err := lib.Load(*pname)
	if err != nil {
		panic(err)
	}

	fmt.Println(fn())
}
