package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandBase16(t *testing.T) {
	var validBase16 = regexp.MustCompile(`^[a-z0-9]+$`)
	testCases := []int{40, 100}
	for _, tcLen := range testCases {
		t.Run(fmt.Sprintf("With len=%d", tcLen), func(t *testing.T) {
			occurences := make(map[string]int)
			for i := 0; i < 10; i++ {
				r := RandBase16(tcLen)
				if len(r) != tcLen {
					t.Errorf("Expected to get string with len=%d, but got len=%d\n", tcLen, len(r))
				}

				if !validBase16.MatchString(r) {
					t.Errorf("Expected to get string with base16 charset, but got '%s'\n", r)
				}
				occurences[r]++
				if occurences[r] != 1 {
					t.Errorf("Expected to get unique string each time, but got duplicate '%s'\n", r)
				}
			}
		})
	}
}

func TestRandBase16Special(t *testing.T) {
	testCases := []int{0, -100}
	for _, tcLen := range testCases {
		t.Run(fmt.Sprintf("With len=%d", tcLen), func(t *testing.T) {
			r := RandBase16(tcLen)
			if r != "" {
				t.Errorf("Expected empty string, but got '%s'\n", r)
			}
		})
	}

}
