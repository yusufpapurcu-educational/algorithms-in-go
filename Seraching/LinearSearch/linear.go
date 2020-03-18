package linearsearch

func LinearSearch(array []int, wanted int) int {
	for i := range array {
		if array[i] == wanted {
			return i
		}
	}
	return -1
}
