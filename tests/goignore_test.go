package tests

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/dev-addict/goignore"
)

type ContentRulesSet struct {
	content string
	rules   goignore.Rules
}

var contentRulesSet = ContentRulesSet{
	"foo\nbar",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar"},
	},
}

var contentRulesSetWithComment = ContentRulesSet{
	"foo\n# comment\nbar",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar"},
	},
}

var contentRulesSetWithEmptyLine = ContentRulesSet{
	"foo\n\n# comment\nbar",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar"},
	},
}

var contentRulesSetWithDirectory = ContentRulesSet{
	"foo\nbar\n\n# comment\nbaz/",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar"},
		{Raw: "baz/", IsDir: true},
	},
}

var contentRulesSetWithNegate = ContentRulesSet{
	"foo\n!bar\n\n# comment\nbaz/",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar", IsNegate: true},
		{Raw: "baz/", IsDir: true},
	},
}

var contentRulesSetWithNegateDirectory = ContentRulesSet{
	"foo\n!bar\n\n# comment\n!baz/",
	goignore.Rules{
		{Raw: "foo"},
		{Raw: "bar", IsNegate: true},
		{Raw: "baz/", IsNegate: true, IsDir: true},
	},
}

func TestParse(t *testing.T) {
	t.Run("should parse valid content", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSet.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSet.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with comment", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSetWithComment.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithComment.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with empty line", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSetWithEmptyLine.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithEmptyLine.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with directory", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSetWithDirectory.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithDirectory.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with negate", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSetWithNegate.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithNegate.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with negate directory", func(t *testing.T) {
		parsedRules, err := goignore.Parse(contentRulesSetWithNegateDirectory.content)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithNegateDirectory.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should not parse content with double star syntax", func(t *testing.T) {
		_, err := goignore.Parse("foo\nfoo/**/bar\nbar")
		if err == nil {
			t.Errorf("Expected error, got nil")
		} else if !errors.Is(goignore.ErrDoubleStarSyntax, err) {
			t.Errorf("Expected error %s, got %s", goignore.ErrDoubleStarSyntax.Error(), err.Error())
		}
	})

	t.Run("should not parse content with invalid rules", func(t *testing.T) {
		_, err := goignore.Parse("# comment\nfoo\nbar\n[123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		} else if !errors.Is(filepath.ErrBadPattern, err) {
			t.Errorf("Expected error %s, got %s", filepath.ErrBadPattern.Error(), err.Error())
		}
	})
}

func TestParseFile(t *testing.T) {
	t.Run("should parse valid content", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSet.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with comment", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-comment")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithComment.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with empty line", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-empty-line")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithEmptyLine.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with directory", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-directory")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithDirectory.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with negate", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-negate")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithNegate.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should parse valid content with negate directory", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-negate-directory")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		parsedRules, err := goignore.ParseFile(file)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSetWithNegateDirectory.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})

	t.Run("should not parse content with double star syntax", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-double-star")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		_, err = goignore.ParseFile(file)
		if err == nil {
			t.Errorf("Expected error, got nil")
		} else if !errors.Is(goignore.ErrDoubleStarSyntax, err) {
			t.Errorf("Expected error %s, got %s", goignore.ErrDoubleStarSyntax.Error(), err.Error())
		}
	})

	t.Run("should not parse content with invalid rules", func(t *testing.T) {
		file, err := os.Open("./testdata/.contentignore-invalid-rule")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		defer file.Close()

		_, err = goignore.ParseFile(file)
		if err == nil {
			t.Errorf("Expected error, got nil")
		} else if !errors.Is(filepath.ErrBadPattern, err) {
			t.Errorf("Expected error %s, got %s", filepath.ErrBadPattern.Error(), err.Error())
		}
	})
}

func TestParseFileFromPath(t *testing.T) {
	t.Run("should parse valid content", func(t *testing.T) {
		parsedRules, err := goignore.ParseFileFromPath("./testdata/.contentignore")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(contentRulesSet.rules, *parsedRules) {
			t.Errorf("Content invalidly parsed")
		}
	})
}
