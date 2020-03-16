package jumpsearch

import (
	"math"
)

func JumpSearch(arr []int, wanted int) int {
	step := int(math.Sqrt(float64(len(arr))))
	onThe := 0
	for range arr {
		if arr[onThe] < wanted {
			onThe += step
			if onThe > len(arr)-1 {
				onThe = len(arr) - 1
			}
		} else if arr[onThe] == wanted {
			return onThe
		} else if arr[onThe] > wanted {
			result := LinearSearch(arr[onThe-step:onThe], wanted)
			if result != -1 {
				return result + onThe - step
			}
			return -1

		}
	}
	return -1
}

func LinearSearch(array []int, wanted int) int {
	for i := range array {
		if array[i] == wanted {
			return i
		}
	}
	return -1
}
