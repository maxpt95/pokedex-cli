package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"split":       {input: "hello world", expected: []string{"hello", "world"}},
		"clean-split": {input: " Should Be LOWERCASE and TRIMMED ", expected: []string{"should", "be", "lowercase", "and", "trimmed"}},
	}

	for name, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("error test [%s]: result length [%d] != expected length [%d]", name, len(result), len(c.expected))
			continue
		}

		for i, word := range result {
			if word != c.expected[i] {
				t.Errorf("error test [%s]: result word %d [%s] != expected [%s]", name, i, word, c.expected[i])
			}
		}
	}
}
