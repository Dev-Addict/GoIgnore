package goignore

import (
	"errors"
	"path/filepath"
)

var (
	ErrDoubleStarSyntax = errors.New("double star syntax is not supported")
	ErrBadPattern       = filepath.ErrBadPattern
)
