package goignore

import (
	"path/filepath"
	"strings"
)

type Rules []Pattern

func (r *Rules) ParseLine(line string) error {
	rule := strings.TrimSpace(line)

	if rule == "" {
		return nil
	}

	if strings.HasPrefix(rule, "#") {
		return nil
	}

	if strings.Contains(rule, "**") {
		return ErrDoubleStarSyntax
	}

	if _, err := filepath.Match(rule, "test"); err != nil {
		return err
	}

	pattern := Pattern{
		Raw: rule,
	}

	if strings.HasPrefix(rule, "!") {
		pattern.IsNegate = true
		pattern.Raw = strings.TrimPrefix(rule, "!")
	}

	if strings.HasSuffix(rule, "/") {
		pattern.IsDir = true
	}

	*r = append(*r, pattern)
	return nil
}

func (r *Rules) Match(path string) (bool, error) {
	for _, rule := range *r {
		matched, err := rule.Match(path)
		if err != nil {
			return false, err
		}
		if matched {
			return !rule.IsNegate, nil
		}
	}

	return false, nil
}
