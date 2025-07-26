// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package util

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

type InputReader struct {
	Stdin    io.Reader
	ReadFile func(string) ([]byte, error)
}

func NewInputReader() *InputReader {
	return &InputReader{
		Stdin:    os.Stdin,
		ReadFile: os.ReadFile,
	}
}

func (r *InputReader) Read(input string) ([]byte, error) {
	switch {
	case input == "-":
		return io.ReadAll(r.Stdin)
	case filepath.Ext(input) == ".json":
		return r.ReadFile(input)
	case input == "":
		return nil, errors.New("empty input")
	default:
		return []byte(input), nil
	}
}
