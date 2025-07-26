// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package marshall

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Captures the output of a function that writes to stdout.
func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	_, _ = buf.ReadFrom(r)
	os.Stdout = stdout
	return buf.String()
}

func TestMarshall(t *testing.T) {
	tests := map[string]struct {
		input     []byte
		expected  string
		expectErr bool
	}{
		"string": {
			input:    []byte(`{"name":"Alice","company":"Company 123"}`),
			expected: `{"company":{"S":"Company 123"},"name":{"S":"Alice"}}`,
		},
		"number": {
			input:    []byte(`{"age":30,"salary":50000.5}`),
			expected: `{"age":{"N":"30"},"salary":{"N":"50000.5"}}`,
		},
		"booelan": {
			input:    []byte(`{"isActive":true,"isAdmin":false}`),
			expected: `{"isActive":{"BOOL":true},"isAdmin":{"BOOL":false}}`,
		},
		"array": {
			input:    []byte(`{"scores":[1.2,3.4],"tags":["go","aws"]}`),
			expected: `{"scores":{"L":[{"N":"1.2"},{"N":"3.4"}]},"tags":{"L":[{"S":"go"},{"S":"aws"}]}}`,
		},
		"map": {
			input:    []byte(`{"details":{"count":10,"items":["item1","item2"]},"meta":{"created":"2025-01-01","valid":true}}`),
			expected: `{"details":{"M":{"count":{"N":"10"},"items":{"L":[{"S":"item1"},{"S":"item2"}]}}},"meta":{"M":{"created":{"S":"2025-01-01"},"valid":{"BOOL":true}}}}`,
		},
		"invalid input": {
			input:     nil,
			expectErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var err error
			output := captureOutput(func() {
				err = Marshall(test.input)
			})

			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.expected+"\n", output)
				assert.NoError(t, err)
			}
		})
	}
}
