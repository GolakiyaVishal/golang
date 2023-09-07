package main

import "testing"

func TestAdd(t *testing.T) {
	// result := Add(2, 3)
	// expected := 5
	// if result != expected {
	// 	t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	// }

	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)

		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d, expected = %d", tc.a, tc.b, result, tc.expected)
		}

	}
}
