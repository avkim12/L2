package main

import "testing"

func TestUnpack(t *testing.T) {

	testTable := []struct {
		str string
		expected string
	} {
		{
			str: "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			str: "abcd",
			expected: "abcd",
		},
		{
			str: "45",
			expected: "",
		},
		{
			str: "",
			expected: "",
		},
		{
			str: "qwe\\4\\5",
			expected: "qwe45",
		},
		{
			str: "qwe\\45",
			expected: "qwe44444",
		},
		{
			str: "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
		},
	}

	for _, testCase := range testTable {
		output, _ := Unpack(testCase.str)
		if output != testCase.expected {
			t.Errorf("Output %q not equal to expected %q", output, testCase.expected)
		}
	}
}