package main

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero")
	}

	return a / b, nil
}
