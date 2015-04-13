package inf

import (
	"errors"
)

func NewDecFromString(s string) (*Dec, error) {
	d, ok := NewDec(0, 0).SetString(s)
	if !ok {
		return nil, errors.New("Bad Decimal string '" + s + "'")
	}
	return d, nil
}
