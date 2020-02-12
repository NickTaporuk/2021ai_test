package operators

import (
	"errors"

	"github.com/NickTaporuk/2021ai_test/set"
)

const (
	// UnionOperatorLex
	UnionOperatorLex = "SUM"
	// DifferentOperatorLex
	DifferentOperatorLex = "DIF"
	// IntersectionOperatorLex
	IntersectionOperatorLex = "INT"
)

var (
	// ErrorFewElements use if the app has not enough elements inside some set
	ErrorFewElements = errors.New(" few elements inside set")
	union            = &Operator{
		Name:          UnionOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) (set.Interface, error) {

			if len(sets) < 2 {
				return nil, ErrorFewElements
			}
			a := sets[0]
			b := sets[1]
			c := sets[2:]
			union := set.Union(a, b, c...)

			return union, nil
		},
	}

	dif = &Operator{
		Name:          DifferentOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) (set.Interface, error) {
			if len(sets) < 2 {
				return nil, ErrorFewElements
			}
			a := sets[0]
			b := sets[1]
			c := sets[2:]

			dif := set.Difference(a, b, c...)

			return dif, nil
		},
	}

	intersect = &Operator{
		Name:          IntersectionOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) (set.Interface, error) {
			if len(sets) < 2 {
				return nil, ErrorFewElements
			}
			a := sets[0]
			b := sets[1]
			c := sets[2:]

			dif := set.Intersection(a, b, c...)

			return dif, nil
		},
	}
)

// nolint
func init() {
	Register(union)
	Register(dif)
	Register(intersect)
}
