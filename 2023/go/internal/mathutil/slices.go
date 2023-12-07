package mathutil

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func Product(arr []int) int {
	prod := arr[0]

	for i := 1; i < len(arr); i++ {
		prod *= arr[i]
	}
	return prod
}
