package goignore

import (
	"path/filepath"
	"strings"
)

type Pattern struct {
	Raw      string
	IsNegate bool
	IsDir    bool
}

func (p *Pattern) Match(path string) (bool, error) {
	if strings.HasPrefix(p.Raw, "/") {
		return filepath.Match(p.Raw, strings.TrimPrefix(path, "/"))
	}
	if strings.Contains(p.Raw, "/") {
		return filepath.Match(p.Raw, path)
	}
	return filepath.Match(p.Raw, filepath.Base(path))
}
