package goignore

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func Parse(content string) (*Rules, error) {
	rules := Rules{}

	for _, line := range strings.Split(content, "\n") {
		if err := rules.ParseLine(line); err != nil {
			return nil, err
		}
	}

	return &rules, nil
}

func ParseFile(file io.Reader) (*Rules, error) {
	rules := Rules{}

	s := bufio.NewScanner(file)
	currentLine := 0
	utf8bom := []byte{0xEF, 0xBB, 0xBF}

	for s.Scan() {
		scannedBytes := s.Bytes()

		if currentLine == 0 {
			scannedBytes = bytes.TrimPrefix(scannedBytes, utf8bom)
		}

		if err := rules.ParseLine(string(scannedBytes)); err != nil {
			return nil, err
		}

		currentLine++
	}

	return &rules, s.Err()
}

func ParseFileFromPath(path string) (*Rules, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ParseFile(file)
}
