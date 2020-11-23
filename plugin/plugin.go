package main

import "github.com/ajm188/go-plugin-repro/lib"

func f() lib.X {
	return 42
}

var F lib.Fn = f
