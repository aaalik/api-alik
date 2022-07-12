package autils

func IsIntSliceContains(slice []int, number int) bool {
	for _, each := range slice {
		if each == number {
			return true
		}
	}

	return false
}

func sumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func sumFloats(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}
