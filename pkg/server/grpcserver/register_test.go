package grpcserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUsername(t *testing.T) {
	ass := assert.New(t)

	var cases = []struct {
		userName string
		expected bool
	}{
		{"a", false},
		{"aBcD", true},
		{"1234567890abcdef", true},
		{"username", true},
		{"_user", false},
		{"user_", false},
		{"user123$", false},
		{"user@domain", false},
		{"", false},
	}

	for _, c := range cases {
		ass.Equal(c.expected, checkUserName(c.userName))
	}
}
