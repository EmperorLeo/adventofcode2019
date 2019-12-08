package util

import "errors"

/*RuneToInt - converts a rune that is basically a single digit into an int */
func RuneToInt(r rune) (int, error) {
	i := int(r - '0')
	if i > 10 || i < 0 {
		return 0, errors.New("r must be a digit from 0-9")
	}
	return i, nil
}
