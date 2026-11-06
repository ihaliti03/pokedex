package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanInput (t *testing.T){
	type testCase struct {
		input string
		expected []string
	}

	testCases := []testCase{
		{
			input: "Hello Wuacapp n0t gonna !ie igot $",
			expected: []string{"hello", "wuacapp", "n0t", "gonna", "!ie", "igot", "$"},
		},
		{
			input: "hello ? ",
			expected: []string{"hello", "?"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "    ",
			expected: []string{},
		},
		{
		    input:    "\thello\tworld\n",
		    expected: []string{"hello", "world"}, 
		},
	}

	for _, test := range testCases{
		got := cleanInput(test.input)
		if !reflect.DeepEqual(test.expected, got){
			t.Errorf(`---------
Test failed
Expecting:	%#v			
Actual:		%#v
Fail			
`, test.expected, got)
		} else {
			fmt.Printf(`---------
Test passed
Expecting:	%#v			
Actual:		%#v
Fail			
`, test.expected, got)
		}
	}
}