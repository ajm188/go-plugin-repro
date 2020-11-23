package lib

import (
	"fmt"
	"plugin"
)

type X int

func Load(path string) (func() X, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup("F")
	if err != nil {
		return nil, err
	}

	fn, ok := f.(func() X)
	if !ok {
		return nil, fmt.Errorf("wrong type for %+v", f)
	}

	return fn, nil
}
