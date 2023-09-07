package main

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b, expected int
		shouldError    bool
	}{
		{4, 2, 2, false},
		{10, 2, 5, false},
		{5, 0, 0, true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)

		if tc.shouldError {
			if err == nil {
				t.Errorf("Divide(%d, %d) should have result in error, but got nil", tc.a, tc.b)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%d, %d) result in unexpected error: %s", tc.a, tc.b, err)
			}

			if result != tc.expected {
				t.Errorf("Divide(%d, %d) = %d; expected: %d", tc.a, tc.b, result, tc.expected)
			}
		}
	}
}
