package operators

import "github.com/NickTaporuk/2021ai_test/set"

var Ops = map[string]*Operator{}

const (
	L = 0
)

// Operator initialize new operator and add logic of operation for this operator
type Operator struct {
	Name          string
	Precedence    int
	Associativity int
	Args          int
	Operation     func(args ...set.Interface) (set.Interface, error)
}

// Register help to register additional operator
func Register(op *Operator) {
	Ops[op.Name] = op
}

// IsOperator check has the operator struct some operator
func IsOperator(str string) bool {
	_, exist := Ops[str]
	return exist
}

// FindOperatorFromString check has struct operator particular operator
func FindOperatorFromString(str string) *Operator {
	op, exist := Ops[str]
	if exist {
		return op
	}
	return nil
}
