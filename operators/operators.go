package operators

import "github.com/fatih/set"

var Ops = map[string]*Operator{}

const (
	L = 0
)

type Operator struct {
	Name          string
	Precedence    int
	Associativity int
	Args          int
	Operation     func(args... set.Interface) set.Interface
}

func Register(op *Operator) {
	Ops[op.Name] = op
}

func IsOperator(str string) bool {
	_, exist := Ops[str]
	return exist
}

func FindOperatorFromString(str string) *Operator {
	op, exist := Ops[str]
	if exist {
		return op
	}
	return nil
}
