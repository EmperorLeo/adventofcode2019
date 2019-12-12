package util

/*Abs - why does the math package not have an abs func for ints??*/
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
