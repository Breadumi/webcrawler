package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	pages := make(map[string]int)

	pages["a"] = 1
	pages["b"] = 2
	pages["c"] = 5
	pages["d"] = 83294
	pages["e"] = 4
	pages["f"] = 8

	tests := []struct {
		name     string
		cfg      *config
		expected []string
	}{
		{
			cfg:      &config{pages: pages},
			name:     "only test",
			expected: []string{"d", "f", "c", "e", "b", "a"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.cfg.sortPages()

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected sort: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
