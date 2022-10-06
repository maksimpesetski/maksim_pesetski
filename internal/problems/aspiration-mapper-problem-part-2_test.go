package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipSting_String(t *testing.T) {
	tests := map[string]struct {
		input, result string
		step          int
	}{
		"Aspiration.com": {
			step:   3,
			input:  "Aspiration.com",
			result: "Aspiration.com",
		},
		"empty sting": {
			step:   3,
			input:  "",
			result: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			processor := NewSkipString(tc.step, tc.input)

			result := processor.String()
			assert.Equal(t, tc.result, result)
		})
	}
}

func TestSkipSting_GetValueAsRuneSlice(t *testing.T) {
	tests := map[string]struct {
		input  string
		step   int
		result []rune
	}{
		"aspiration.com": {
			step:   3,
			input:  "Aspiration.com",
			result: []rune("Aspiration.com"),
		},
		"HaPPyBirthday2you!": {
			step:   3,
			input:  "HaPPyBirthday2you!",
			result: []rune("HaPPyBirthday2you!"),
		},
		"May be 2morrow": {
			step:   3,
			input:  "May be 2morrow",
			result: []rune("May be 2morrow"),
		},
		"empty sting": {
			step:   3,
			input:  "",
			result: []rune(""),
		},
		"upper case sting": {
			step:   3,
			input:  "HELLO WORLD",
			result: []rune("HELLO WORLD"),
		},
		"lower case sting": {
			step:   3,
			input:  "hello world",
			result: []rune("hello world"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			processor := NewSkipString(tc.step, tc.input)

			result := processor.GetValueAsRuneSlice()

			assert.Equal(t, tc.result, result)
		})
	}
}

func TestSkipSting_TransformRune(t *testing.T) {
	tests := map[string]struct {
		s         string
		step, pos int
		result    SkipSting
	}{
		"first char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  0,
			result: SkipSting{
				value: []rune("aspiration.com"),
				step:  3,
			},
		},
		"second char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  1,
			result: SkipSting{
				value: []rune("Aspiration.com"),
				step:  3,
			},
		},
		"third char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  2,
			result: SkipSting{
				value: []rune("AsPiration.com"),
				step:  3,
			},
		},
		". char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  10,
			result: SkipSting{
				value: []rune("Aspiration.com"),
				step:  3,
			},
		},
		"12th char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  11,
			result: SkipSting{
				value: []rune("Aspiration.com"),
				step:  3,
			},
		},
		"13th char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  12,
			result: SkipSting{
				value: []rune("Aspiration.cOm"),
				step:  3,
			},
		},
		"6th char with step 3": {
			step: 3,
			s:    "Aspiration.com",
			pos:  5,
			result: SkipSting{
				value: []rune("AspirAtion.com"),
				step:  3,
			},
		},
		"out of boundary position sting": {
			step: 3,
			s:    "Aspiration.com",
			pos:  100,
			result: SkipSting{
				value: []rune("Aspiration.com"),
				step:  3,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			processor := NewSkipString(tc.step, tc.s)

			processor.TransformRune(tc.pos)

			assert.Equal(t, tc.result.String(), processor.String())
		})
	}
}
