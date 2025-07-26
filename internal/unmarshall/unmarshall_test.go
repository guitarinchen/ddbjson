// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package unmarshall

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

func TestUnmarshall(t *testing.T) {
	tests := map[string]struct {
		input     []byte
		expected  string
		expectErr bool
	}{
		"valid input": {
			input:    []byte(`{"companyId":{"S":"COMP123"},"userId":{"S":"USER456"},"address":{"M":{"city":{"S":"Anytown"},"state":{"S":"CA"},"street":{"S":"123 Main St"},"zipCode":{"S":"12345"}}},"age":{"N":"30"},"departments":{"L":[{"S":"HR"},{"S":"Finance"}]},"email":{"S":"john.doe@example.com"},"firstName":{"S":"John"},"isActive":{"BOOL":true},"lastLogin":{"S":"2023-05-15T14:30:00Z"},"lastName":{"S":"Doe"},"projects":{"L":[{"S":"Project A"},{"S":"Project B"},{"S":"Project C"}]},"salary":{"N":"75000.5"}}`),
			expected: `{"address":{"city":"Anytown","state":"CA","street":"123 Main St","zipCode":"12345"},"age":30,"companyId":"COMP123","departments":["HR","Finance"],"email":"john.doe@example.com","firstName":"John","isActive":true,"lastLogin":"2023-05-15T14:30:00Z","lastName":"Doe","projects":["Project A","Project B","Project C"],"salary":75000.5,"userId":"USER456"}`,
		},
		"invalid input": {
			input:     []byte("invalid"),
			expectErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var err error
			output := captureOutput(func() {
				err = Unmarshall(test.input)
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
