package inf

import (
	"errors"
)

func Parse(s string) (*Dec, error) {
	d, ok := NewDec(0, 0).SetString(s)
	if !ok {
		return nil, errors.New("Bad Decimal string '" + s + "'")
	}
	return d, nil
}

func MustParse(s string) *Dec {
	d, ok := NewDec(0, 0).SetString(s)
	if !ok {
		panic("Bad Decimal string '" + s + "'")
	}
	return d
}
