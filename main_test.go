package main

import "testing"

type  TestCase struct {
	isDecode bool
	input string
	expected string
}

func TestEncodeOrDecode(t *testing.T) {
	testCases := []TestCase {
		{false, "hello", "NBSWY3DP"},
		{true, "NBSWY3DP", "hello"},
	}
	for _, testCase := range testCases {
		// given
		input := testCase.input
		expected := testCase.expected
		isDecode := testCase.isDecode

		//when
		var actual []byte
		if isDecode {
			actual = decode([]byte(input))
		} else {
			actual = encode([]byte(input))
		}


		// then
		if string(actual) != expected {
			t.Errorf("Expected: %s, got: %s", expected, string(actual))
		}
	}
}

func TestEncodeOrDecodeWithAnonStruct(t *testing.T) {
	testCases := []struct{
		isDecode bool
		input string
		expected string
	}{
		{false, "hello", "NBSWY3DP"},
		{true, "NBSWY3DP", "hello"},
	}

	for _, testCase := range testCases {
		// given
		input := testCase.input
		expected := testCase.expected
		isDecode := testCase.isDecode

		//when
		var actual []byte
		if isDecode {
			actual = decode([]byte(input))
		} else {
			actual = encode([]byte(input))
		}

		// then
		if string(actual) != expected {
			t.Errorf("Expected: %s, got: %s", expected, string(actual))
		}
	}
}



