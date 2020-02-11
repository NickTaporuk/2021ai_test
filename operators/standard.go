package operators

import (
	"github.com/fatih/set"
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
	union = &Operator{
		Name:          UnionOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) set.Interface {

			a := sets[0]
			b := sets[1]
			c := sets[2:]
			union := set.Union(a, b, c...)

			return union
		},
	}

	dif = &Operator{
		Name:          DifferentOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) set.Interface {
			a := sets[0]
			b := sets[1]
			c := sets[2:]

			dif := set.Difference(a, b, c...)

			return dif
		},
	}

	intersect = &Operator{
		Name:          IntersectionOperatorLex,
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(sets ...set.Interface) set.Interface {

			a := sets[0]
			b := sets[1]
			c := sets[2:]

			dif := set.Intersection(a, b, c...)

			return dif
		},
	}
)

func init() {
	Register(union)
	Register(dif)
	Register(intersect)
}
