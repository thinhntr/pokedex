package cmd

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello     world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    help     mapb map   ",
			expected: []string{"help", "mapb", "map"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "goodbye ",
			expected: []string{"goodbye"},
		},
		{
			input:    "    run Super  At0M1c  Poke",
			expected: []string{"run", "super", "at0m1c", "poke"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("len(actual)=%d != len(c.expected)=%d", len(actual), len(c.expected))
			continue
		}

		i := 0
		for i = range actual {
			if actual[i] != c.expected[i] {
				i -= 1
				break
			}
		}
		if i < len(c.expected)-1 {
			t.Errorf("not match.\nActual: %v\nExpected:%v", actual, c.expected)
		}
	}
}
