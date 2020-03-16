package LinearSerach

import "fmt"

func main() {
	var exampleArray []int = []int{10, 2, 5, 13, 65, 97, 354, 45, 8}
	fmt.Println(LinearSerach(exampleArray, 97))
}

func LinearSerach(array []int, wanted int) int {
	for i := range array {
		if array[i] == wanted {
			return i
		}
	}
	return -1
}
