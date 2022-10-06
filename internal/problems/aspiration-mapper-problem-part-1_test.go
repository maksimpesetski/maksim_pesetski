package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	tests := map[string]struct {
		input, result string
	}{
		"aspiration.com": {
			input:  "Aspiration.com",
			result: "asPirAtiOn.cOm",
		},
		"HaPPyBirthday2you!": {
			input:  "HaPPyBirthday2you!",
			result: "haPpyBirThdAy2yOu!",
		},
		"May be 2morrow": {
			input:  "May be 2morrow",
			result: "maY be 2MorRow",
		},
		"upper case sting": {
			input:  "HELLO WORLD",
			result: "heLlo WorLd",
		},
		"lower case sting": {
			input:  "hello world",
			result: "heLlo WorLd",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			result := CapitalizeEveryThirdAlphanumericChar(tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}
