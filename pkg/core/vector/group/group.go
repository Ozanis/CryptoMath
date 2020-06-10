package group

import (
	"cryptops/pkg/core/vector/operator"
)

type Grouper interface {
	Build(mod operator.Vector)
	Operate(dest, x, y operator.Vector) operator.Vector
	Size(x, y operator.Vector) uint8
}
