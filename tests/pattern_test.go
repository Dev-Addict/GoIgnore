package tests

import (
	"testing"

	"github.com/dev-addict/goignore"
)

func TestPattern(t *testing.T) {
	t.Parallel()

	t.Run("Match", func(t *testing.T) {
		t.Run("Should match", func(t *testing.T) {
			pattern := goignore.Pattern{
				Raw: "foo",
			}

			matched, err := pattern.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}
		})

		t.Run("Should match directory", func(t *testing.T) {
			pattern := goignore.Pattern{
				Raw:   "foo/",
				IsDir: true,
			}

			matched, err := pattern.Match("foo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("bar/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}

			matched, err = pattern.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}
		})

		t.Run("Should match wildcard", func(t *testing.T) {
			pattern := goignore.Pattern{
				Raw: "foo*",
			}

			matched, err := pattern.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("foobar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}

			matched, err = pattern.Match("barfoo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}
		})

		t.Run("Should match wildcard directory", func(t *testing.T) {
			pattern := goignore.Pattern{
				Raw:   "foo*/",
				IsDir: true,
			}

			matched, err := pattern.Match("foo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("foobar/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("bar/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}

			matched, err = pattern.Match("barfoo/")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}

			matched, err = pattern.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}
		})

		t.Run("Should match without negation consideration", func(t *testing.T) {
			pattern := goignore.Pattern{
				Raw:      "foo",
				IsNegate: true,
			}

			matched, err := pattern.Match("foo")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if !matched {
				t.Errorf("Should match")
			}

			matched, err = pattern.Match("bar")
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			} else if matched {
				t.Errorf("Should not match")
			}
		})
	})
}
