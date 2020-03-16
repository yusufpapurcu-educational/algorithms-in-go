package jumpsearch

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestJumpSearch(t *testing.T) {
	a := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610} // Test Dizisi
	x := a[rand.Int()%len(a)]                                               // Istenen eleman

	i := JumpSearch(a, x) // Arama Algoritması
	if i < len(a) && a[i] == x {
		//		fmt.Printf("%d Sayısı dizinin %d. Indexinde bulundu. Dizi :  %v\n", x, i, a)
	} else {
		t.Fatalf("%d Sayısı, %v Dizisi içinde bulunamadı\n", x, a)
	}
}

func BenchmarkJumpSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610} // Test Dizisi
		x := a[rand.Int()%len(a)]                                               // Istenen eleman

		i := JumpSearch(a, x) // Arama Algoritması
		if i < len(a) && a[i] == x {
			fmt.Printf("%d Sayısı dizinin %d. Indexinde bulundu. Dizi :  %v\n", x, i, a)
		} else {
			log.Fatalf("%d Sayısı, %v Dizisi içinde bulunamadı\n", x, a)
		}
	}
}
