package group

import "cryptops/pkg/core/vector/operator"

type Sum struct {
	Weight Vector
	Mod    Vector
}

func (Sum) Build(mod operator.Vector) operator.Vector {

}

func (Sum) Operate(dest, x, y operator.Vector) operator.Vector {
	return dest
}
