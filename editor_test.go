package main

import "testing"

func TestGetStartOfLine(t *testing.T) {
	cases := []struct {
		gapBuffer GapBuffer
		input     int
		expected  int
	}{
		{
			gapBuffer: GapBuffer{
				text:     []rune{},
				gapStart: 0,
				gapEnd:   0,
			},
			input:    0,
			expected: 0,
		},
		{
			gapBuffer: GapBuffer{
				text:     []rune{'h', 'i', rune(0), rune(0), rune(0)},
				gapStart: 2,
				gapEnd:   6,
			},
			input:    1,
			expected: 0,
		},
		{
			gapBuffer: GapBuffer{
				text:     []rune{'h', 'e', 'l', 'l', 'o', rune(0), rune(0), rune(0)},
				gapStart: 5,
				gapEnd:   8,
			},
			input:    6,
			expected: 0,
		},
		{
			gapBuffer: GapBuffer{
				text:     []rune{'h', 'i', '\r', 'y', 'o', rune(0), rune(0), rune(0), '!'},
				gapStart: 5,
				gapEnd:   8,
			},
			input:    5,
			expected: 3,
		},
		{
			gapBuffer: GapBuffer{
				text:     []rune{'\r', rune(0), rune(0), rune(0)},
				gapStart: 1,
				gapEnd:   4,
			},
			input:    2,
			expected: 1,
		},
	}

	for _, c := range cases {
		output := c.gapBuffer.getStartOfLine(c.input)
		if output != c.expected {
			t.Errorf("Expected %d, got %d", output, c.expected)
		}
	}
}
