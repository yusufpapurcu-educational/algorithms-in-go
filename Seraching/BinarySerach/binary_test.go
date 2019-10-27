package binaryserach

import (
	"fmt"
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55} // Test Dizisi
	x := 45                                         // Istenen eleman

	i := Search(len(a), func(i int) bool { return a[i] >= x }) // Arama Algoritması
	if i < len(a) && a[i] == x {
		fmt.Printf("%d Sayısı dizinin %d. Indexinde bulundu. Dizi :  %v\n", x, i, a)
	} else {
		t.Fatalf("%d Sayısı, %v Dizisi içinde bulunamadı\n", x, a)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55} // Test Dizisi
		//	x := a[rand.Int()%len(a)]                       // Istenen eleman
		x := 6
		i := Search(len(a), func(i int) bool { return a[i] >= x }) // Arama Algoritması
		if i < len(a) && a[i] == x {
			//	fmt.Printf("%d Sayısı dizinin %d. Indexinde bulundu. Dizi :  %v\n", x, i, a)
		} else {
			log.Fatalf("%d Sayısı, %v Dizisi içinde bulunamadı\n", x, a)
		}
	}
}
