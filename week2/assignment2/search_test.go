package main

import "testing"

func TestMatchesSearchPattern(t *testing.T) {
	var tests = []struct {
		s     string
		found bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			matchesSearchPattern := matchesSearchPattern(test.s)
			if matchesSearchPattern != test.found {
				t.Errorf("Error with '%v', expecting '%v', actual '%v'", test.s, test.found, matchesSearchPattern)
			}
		})
	}
}
