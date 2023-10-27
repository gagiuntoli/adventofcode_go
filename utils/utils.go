package utils

func ArrayMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func ArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
