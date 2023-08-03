# GoIgnore

A simple tool to work with .ignore files in Go.

## Installation

```bash
go get github.com/monochromegane/goignore
```

## Documentation

- [Functions](#functions)
    - [Parse](#parse)
    - [ParseFile](#parsefile)
    - [ParseFileFromPath](#parsefilefrompath)
- [Types](#types)
    - [Pattern](#pattern)
        - [Match](#patternmatch)
    - [Rules](#rules)
        - [ParseLine](#rulesparseline)
- [Errors](#errors)
    - [ErrDoubleStarSyntax](#errdoublestarsyntax)
    - [ErrBadPattern](#errbadpattern)

### Functions

#### Parse

Parse parses the given ignore file content and returns the Rules.

```go
func Parse(content string) (*Rules, error)
```

Example:

```go
rules, err := goignore.Parse("# comment\nfoo\nbar")
if err != nil {
panic(err)
}

fmt.Println(rules.Match("foo")) // => true
fmt.Println(rules.Match("bar")) // => true
fmt.Println(rules.Match("baz")) // => false
```

#### ParseFile

ParseFile parses the given ignore file and returns the Rules.

```go
func ParseFile(file *os.File) (*Rules, error)
```

Example:

```go
file, err := os.Open("testdata/.contentignore")
if err != nil {
panic(err)
}
defer file.Close()

rules, err := goignore.ParseFile(file)
if err != nil {
panic(err)
}

fmt.Println(rules.Match("foo")) // => true
fmt.Println(rules.Match("bar")) // => true
fmt.Println(rules.Match("baz")) // => false
```

#### ParseFileFromPath

ParseFileFromPath parses the given ignore file path and returns the Rules.

```go
func ParseFileFromPath(path string) (*Rules, error)
```

Example:

```go
rules, err := goignore.ParseFileFromPath("testdata/.contentignore")
if err != nil {
panic(err)
}

fmt.Println(rules.Match("foo")) // => true
fmt.Println(rules.Match("bar")) // => true
fmt.Println(rules.Match("baz")) // => false
```

### Types

#### Pattern

Pattern represents a pattern of ignore file.

```go
type Pattern struct {
Raw string    // Raw is a raw pattern string.
IsNegate bool // IsNegate is a flag that the pattern is negated.
IsDir bool    // IsDir is a flag that the pattern is directory.
}
```

##### Pattern.Match

Match returns true if the given path matches the pattern. IsNegate and IsDir are not considered.

```go
func (p *Pattern) Match(path string) (bool, error)
```

Example:

```go
pattern := &goignore.Pattern{Raw: "foo"}

fmt.Println(pattern.Match("foo")) // => true
fmt.Println(pattern.Match("bar")) // => false
```

#### Rules

Rules represents a set of ignore patterns.

```go
type Rules []Pattern
```

##### Rules.ParseLine

ParseLine parses the given line and appends the pattern to the Rules.

```go
func (r *Rules) ParseLine(line string) error
```

Example:

```go
rules := goignore.Rules{}

err := rules.ParseLine("foo")
if err != nil {
panic(err)
}

fmt.Println(rules.Match("foo")) // => true
fmt.Println(rules.Match("bar")) // => false
```

### Errors

#### ErrDoubleStarSyntax

ErrDoubleStarSyntax is an error that the pattern contains invalid double star syntax.

```go
import "errors"

rules, err := goignore.Parse("foo/**/bar")

fmt.Println(errors.Is(err, ErrDoubleStarSyntax)) // => true
```

#### ErrBadPattern

ErrBadPattern is an error that the pattern is invalid.

```go
import (
    "errors"
	"path/filepath"
)

rules, err := goignore.Parse("[123")

fmt.Println(errors.Is(err, ErrBadPattern)) // => true
fmt.Println(errors.Is(err, filepath.ErrBadPattern)) // => true
```

## Contributing

Simply fork the repository and send a pull request.

## License

[MIT](./LICENSE)
