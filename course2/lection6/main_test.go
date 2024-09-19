package main

import "testing"

var factorialTestCases = []struct {
	input    int
	expected int
}{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{5, 120},
	{6, 720},
}

func TestFactorialTests(t *testing.T) {
	for idx, testCase := range factorialTestCases {
		if res := factorial(testCase.input); res != testCase.expected {
			t.Fatalf("case %d. factorial(%d) returns %d, when %d expected\n", idx, testCase.input, res, testCase.expected)
		}
	}
}
