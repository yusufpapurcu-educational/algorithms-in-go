package quicksort

import (
	"math/rand"
)

// QuickSort Algoritması
func quickSort(dizi []int) []int {
	if len(dizi) < 2 {
		return dizi
	}

	left, right := 0, len(dizi)-1

	// Dayanak Seçimi
	pivot := rand.Int() % len(dizi)

	dizi[pivot], dizi[right] = dizi[right], dizi[pivot]

	for i := range dizi {
		if dizi[i] < dizi[right] {
			dizi[left], dizi[i] = dizi[i], dizi[left]
			left++
		}
	}

	dizi[left], dizi[right] = dizi[right], dizi[left]

	quickSort(dizi[:left])
	quickSort(dizi[left+1:])

	return dizi
}
