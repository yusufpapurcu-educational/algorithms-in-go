package binaryserach

// Search fonksiyonu bizim BinarySerach algoritmamız
func Search(n int, f func(int) bool) int { // Burada dizi uzunluğu ve fonksiyon alma amacı fazladan dizi oluşturmadan gelen fonksiyonu kullanarak büyük ya da küçük olduğunu öğrenmek.

	i, j := 0, n // i degiskeni istenen sayının indexini alacak. Simdilik 0.

	for i < j {
		h := i + (j-i)/2 // Burada aranan elemanı bulmak için kullanacağımız dayank belirleniyor.

		if !f(h) { // Dizinin ortasındaki rakam aranan rakamdan büyük mü, küçük mü o hesaplanıyor.
			i = h + 1 // Eğer aranan eleman dayanaktan büyük ise i yeniden tanımlanıyor.
		} else {
			j = h // Küçük ise üst değer dayanağa eşitleniyor.
		}
	}
	// Döngü bitince elde edilen i değeri aranan sayının indexini verir.
	return i
}
