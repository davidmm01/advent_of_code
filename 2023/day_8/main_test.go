package main

// import (
// 	"fmt"
// 	"testing"
// )

// func TestAllPositionsEndInZ(t *testing.T) {
// 	var tests = []struct {
// 		input    []string
// 		expected bool
// 	}{
// 		{
// 			input:    []string{"AAZ", "AAZ", "ZZZ"},
// 			expected: true,
// 		},
// 		{
// 			input:    []string{"AAA", "AAZ", "ZZZ"},
// 			expected: false,
// 		},
// 		{
// 			input:    []string{"AAA", "AAA", "ZZA"},
// 			expected: false,
// 		},
// 		{
// 			input:    []string{"ZZA", "ZZZ", "ZZZ"},
// 			expected: false,
// 		},
// 	}

// 	for i, tt := range tests {
// 		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
// 			if allPositionsEndInZ(tt.input) != tt.expected {
// 				t.Errorf("tt.input %v, tt.expected %v failed", tt.input, tt.expected)
// 			}
// 		})
// 	}
// }
