// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package util

import (
	"errors"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := map[string]struct {
		input       string
		stdin       string
		readFile    func(string) ([]byte, error)
		expected    string
		expectedErr string
	}{
		"stdin": {
			input:    "-",
			stdin:    `{"from":"stdin"}`,
			expected: `{"from":"stdin"}`,
		},
		"valid json file": {
			input: "test.json",
			readFile: func(path string) ([]byte, error) {
				return []byte(`{"from":"file"}`), nil
			},
			expected: `{"from":"file"}`,
		},
		"invalid json file": {
			input: "invalid.json",
			readFile: func(path string) ([]byte, error) {
				return nil, errors.New("no such file or directory")
			},
			expectedErr: "no such file or directory",
		},
		"default input": {
			input:    `{"raw":"string"}`,
			expected: `{"raw":"string"}`,
		},
		"empty input": {
			input:       "",
			expectedErr: "empty input",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reader := &InputReader{
				Stdin:    strings.NewReader(test.stdin),
				ReadFile: test.readFile,
			}

			res, err := reader.Read(test.input)

			if test.expectedErr != "" {
				if err == nil || !strings.Contains(err.Error(), test.expectedErr) {
					t.Errorf("expected error: %s, got: %v", test.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if string(res) != test.expected {
					t.Errorf("expected: %s, got: %s", test.expected, res)
				}
			}
		})
	}
}
