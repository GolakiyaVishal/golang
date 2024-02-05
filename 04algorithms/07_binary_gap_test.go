package main

import "testing"

func TestFindGap(t *testing.T) {

	testCases := []struct {
		a, expected int
	}{
		{1041, 5},
		{15, 0},
		{32, 0},
	}

	for _, tc := range testCases {
		result := FindGap(int32(tc.a))

		if result != tc.expected {
			t.Errorf("Gap is %d, expected %d", result, tc.expected)
		}
	}

}
