package main

import (
	"fmt"
	"testing"
)

// basic unit testing

//func TestAdd(t *testing.T) {
//	result := Add(1, 2)
//	expected := 3
//
//	if result != expected {
//		t.Errorf("Add(1, 2) failed, got %d, want %d", result, expected)
//	}
//}

// table driven unit testing

type testCase struct {
	a, b     int
	expected int
}

func TestAdd(t *testing.T) {

	tests := []testCase{
		{a: 1, b: 2, expected: 3},
		{a: -1, b: -2, expected: -3},
		{a: 0, b: 5, expected: 5},
		{a: 100, b: 200, expected: 300},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%d,%d", tc.a, tc.b)
		t.Run(testname, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) got %d, want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
