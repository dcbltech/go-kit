package hashutils

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Hex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", fmt.Sprintf("%x", md5.Sum([]byte("hello")))},
		{"world", fmt.Sprintf("%x", md5.Sum([]byte("world")))},
		{"", fmt.Sprintf("%x", md5.Sum([]byte("")))},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := MD5Hex(test.input)
			assert.Equal(t, test.expected, result, "MD5Hex(%q) = %q; want %q", test.input, result, test.expected)
		})
	}
}
