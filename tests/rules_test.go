package tests

import (
	"errors"
	"testing"

	"github.com/dev-addict/goignore"
)

var rules goignore.Rules

func init() {
	rules = goignore.Rules{}
}

func TestRules(t *testing.T) {
	t.Run("ParseLine", func(t *testing.T) {
		rulesLength := len(rules)

		t.Run("should not parse empty line", func(t *testing.T) {
			err := rules.ParseLine("")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			if len(rules) != rulesLength {
				t.Errorf("Rules should be empty")
			}
		})

		t.Run("should not parse comment line", func(t *testing.T) {
			err := rules.ParseLine("# comment")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			if len(rules) != rulesLength {
				t.Errorf("Rules should be empty")
			}
		})

		t.Run("should not parse double star syntax", func(t *testing.T) {
			err := rules.ParseLine("foo/**/bar")
			if !errors.Is(err, goignore.ErrDoubleStarSyntax) {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if err == nil {
				t.Errorf("Error should not be nil")
			}

			if len(rules) != rulesLength {
				t.Errorf("Rules should be empty")
			}
		})

		t.Run("should not parse invalid pattern", func(t *testing.T) {
			err := rules.ParseLine("[")
			if !errors.Is(err, goignore.ErrBadPattern) {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if err == nil {
				t.Errorf("Error should not be nil")
			}

			if len(rules) != rulesLength {
				t.Errorf("Rules should be empty")
			}
		})

		t.Run("should parse pattern", func(t *testing.T) {
			err := rules.ParseLine("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			rulesLength++

			if len(rules) != rulesLength {
				t.Errorf("Rules should have %d item", rulesLength)
			} else if rules[rulesLength-1].Raw != "foo" {
				t.Errorf("Rule should be foo")
			}
		})

		t.Run("should parse negate pattern", func(t *testing.T) {
			err := rules.ParseLine("!foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			rulesLength++

			if len(rules) != rulesLength {
				t.Errorf("Rules should have %d item", rulesLength)
			} else if rules[rulesLength-1].Raw != "foo" {
				t.Errorf("Rule should be foo")
			} else if !rules[rulesLength-1].IsNegate {
				t.Errorf("Rule should be negate")
			}
		})

		t.Run("should parse directory pattern", func(t *testing.T) {
			err := rules.ParseLine("foo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			rulesLength++

			if len(rules) != rulesLength {
				t.Errorf("Rules should have %d item", rulesLength)
			} else if rules[rulesLength-1].Raw != "foo/" {
				t.Errorf("Rule should be foo")
			} else if !rules[rulesLength-1].IsDir {
				t.Errorf("Rule should be directory")
			}
		})

		t.Run("should parse negate directory pattern", func(t *testing.T) {
			err := rules.ParseLine("!foo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			rulesLength++

			if len(rules) != rulesLength {
				t.Errorf("Rules should have %d item", rulesLength)
			} else if rules[rulesLength-1].Raw != "foo/" {
				t.Errorf("Rule should be foo")
			} else if !rules[rulesLength-1].IsDir {
				t.Errorf("Rule should be directory")
			} else if !rules[rulesLength-1].IsNegate {
				t.Errorf("Rule should be negate")
			}
		})

		t.Run("should parse pattern with \"/\" prefix", func(t *testing.T) {
			err := rules.ParseLine("/foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			rulesLength++

			if len(rules) != rulesLength {
				t.Errorf("Rules should have %d item", rulesLength)
			} else if rules[rulesLength-1].Raw != "/foo" {
				t.Errorf("Rule should be /foo")
			}
		})

		t.Cleanup(func() {
			rules = goignore.Rules{}
			rulesLength = len(rules)
		})
	})

	t.Run("Match", func(t *testing.T) {
		t.Run("should not match empty rules", func(t *testing.T) {
			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match foo")
			}
		})

		t.Run("should not match empty path", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "foo"},
			}

			matched, err := rules.Match("")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match empty path")
			}
		})

		t.Run("should not match empty pattern", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: ""},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match foo")
			}
		})

		t.Run("should match exact pattern", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "foo"},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match foo")
			}

			matched, err = rules.Match("bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match bar")
			}
		})

		t.Run("should match pattern directories", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "foo/", IsDir: true},
			}

			matched, err := rules.Match("foo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if !matched {
				t.Errorf("Should match foo")
			}

			matched, err = rules.Match("bar/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match bar")
			}

			matched, err = rules.Match("bar/foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match bar/foo")
			}

			matched, err = rules.Match("bar/foo/bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match bar/foo/bar")
			}
		})

		t.Run("should match negate pattern", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "foo", IsNegate: true},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo")
			}
		})

		t.Run("should match negate pattern with \"/\" prefix", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "/foo", IsNegate: true},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo")
			}
		})

		t.Run("should match negate pattern directories", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "foo/", IsDir: true, IsNegate: true},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo")
			}

			matched, err = rules.Match("foo/bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo/bar")
			}
		})

		t.Run("should match negate pattern with \"/\" prefix directories", func(t *testing.T) {
			rules = goignore.Rules{
				{Raw: "/foo", IsDir: true, IsNegate: true},
			}

			matched, err := rules.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo")
			}

			matched, err = rules.Match("foo/bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if matched {
				t.Errorf("Should not match foo/bar")
			}
		})
	})
}
