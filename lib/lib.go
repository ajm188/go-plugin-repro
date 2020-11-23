package lib

import (
	"fmt"
	"plugin"
)

type X int

type Fn func() X

func Load(path string) (Fn, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup("F")
	if err != nil {
		return nil, err
	}

	fn, ok := f.(*Fn)
	if !ok {
		return nil, fmt.Errorf("wrong type for %+v", f)
	}

	return *fn, nil
}
