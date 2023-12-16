package mathutil

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
