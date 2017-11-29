package main

import (
	"regexp"
	"testing"
)

func TestRandBase16(t *testing.T) {
	var validbase16 = regexp.MustCompile(`^[a-z0-9]+$`)
	occurences := make(map[string]int)
	l := 40

	for i := 0; i < 10; i++ {
		r := RandBase16(l)
		if len(r) != l {
			t.Errorf("Expected to get string on len=%d, but got len=%d\n", l, len(r))
		}

		if !validbase16.MatchString(r) {
			t.Errorf("Expected to get string with base16 charset, but got '%s'\n", r)
		}
		occurences[r]++
		if occurences[r] != 1 {
			t.Errorf("Expected to get unique string each time, but got duplicate '%s'\n", r)
		}
	}
}
