package storage

import (
	"fmt"
)

type Engine interface {
	fmt.Stringer
	Save(entity Entity) error
}
