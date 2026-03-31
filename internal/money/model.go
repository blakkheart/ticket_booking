package money

import (
	"errors"
	"math/big"
)

type Money struct {
	value *big.Rat
}

func NewMoney(s string) (*Money, error) {
	r := new(big.Rat)
	if _, ok := r.SetString(s); !ok {
		return nil, errors.New("Invalid money format")
	}
	if r.Sign() < 0 {
		return nil, errors.New("Negative money")
	}
	return &Money{value: r}, nil
}

func (m *Money) String() string {
	return m.value.FloatString(2)
}

func (m *Money) Add(other *Money) *Money {
	return &Money{
		value: new(big.Rat).Add(m.value, other.value),
	}
}

func (m *Money) GreaterThan(other *Money) bool {
	return m.value.Cmp(other.value) > 0
}
