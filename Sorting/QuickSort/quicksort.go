package quicksort

import (
	"math/rand"
)

// QuickSort Algoritması
func quickSort(dizi []int) []int {
	if len(dizi) < 2 { // Eğer dizi 1 elemanlı ise diziyi dönüyor. Bu fonksiyonun recursive çalışması için önemli.
		return dizi
	}

	left, right := 0, len(dizi)-1 // Dizinin ilk ve son elemanlarının indexi alınıyor.

	rand.Seed(int64(len(dizi))) // Random için besleme yapılıyor. Dayanak seçimi için önemli.

	// Dayanak Seçimi
	pivot := rand.Int() % len(dizi)

	dizi[pivot], dizi[right] = dizi[right], dizi[pivot] // Dayanak ile son elemanın yerini değiştiriyoruz.

	for i := range dizi { // Dizi boyunda bir dönguye giriyoruz.
		if dizi[i] < dizi[right] { // Dizinin i elemanı dayanaktan küçük ise
			dizi[left], dizi[i] = dizi[i], dizi[left] // Onu ilk indexe alıyoruz
			left++                                    // Ve ilk indexi bir arttırıyoruz.
		}
	}

	dizi[left], dizi[right] = dizi[right], dizi[left] // sonrasında ilk index ile dayanağı yer değiştiriyoruz.

	quickSort(dizi[:left])   // ilk yarıyı tekrardan sıralamaya sokuyoruz.
	quickSort(dizi[left+1:]) // ikinci yarı da sıralamaya giriyor.

	// Bu işlemlerin tamamı dizi üzerinde yapıldığı için recursive sonunda dizi sıralanmış oluyor.
	// Böylece gereksiz bellek ve dizi kullanmadan en yüksek performansta tüm dizideki elemanları sıralıyoruz.
	return dizi
}
