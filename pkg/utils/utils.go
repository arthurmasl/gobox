package utils

import "errors"

func Summy(a, b int) (int, error) {
	if a == 555 {
		return -1, errors.New("summy err")
	}
	return a + b, nil
}
