package storage

import (
	"fmt"
	"io"
)

type Engine interface {
	fmt.Stringer
	io.Closer
	Save(entity Entity) error
}
