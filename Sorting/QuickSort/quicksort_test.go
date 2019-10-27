package quicksort

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, test := range tests { // Testler içerisinde dolanıyoruz
		quickSort(test.input)          // Fonksiyona test.input'u verdik. Fonksiyon bitinde test.input sıralanmış olacak.
		for i := range test.expected { // Burada sıralanan ile beklenen dizilerini karşılaştırmak için döngüye girdik.
			if test.input[i] != test.expected[i] { // İsterseniz test.expected listesinden bir değeri bozarak kontrol edi etmediğine bakabilirsiniz.
				t.Errorf("QuickSorted = %d, expected %d.",
					test.input, test.expected)
			}
		}
	}
}

// Tamamen test ile aynı sadece Benchmark
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			quickSort(test.input)
			for i := range test.expected {
				if test.input[i] != test.expected[i] {
					log.Fatalf("Expected : %d, Sorted : %d", test.input, test.expected)
				}
			}
		}
	}
}

//Eğer tests struct dizisindeki değerler doğru olmasına rağmen go test hata veriyorsa algoritmada sorun var demektir.
var tests = []struct { // Test dizimiz. İçinde örnek ve cevaplar var.
	input    []int // Girdi
	expected []int // Beklenen
}{
	{
		input:    []int{9, 4, 7, 1, -6, 8, -3, 5, -9, 2},
		expected: []int{-9, -6, -3, 1, 2, 4, 5, 7, 8, 9},
	},
	{
		input:    []int{10, 11, -14, -9, 4, 10, -4, 8, 10, 21, 2, 10, -3, 2, -29, 11, -7, 8, -11, 26, -25, -27, -8, -11, -38, 10, -13, 17, 37, 12, 1, -3, -15, 26, 13, -3, -1, -14, 16, -16, -5, -39, -16, -24, 31, 15, 5, 10, -18, -18},
		expected: []int{-39, -38, -29, -27, -25, -24, -18, -18, -16, -16, -15, -14, -14, -13, -11, -11, -9, -8, -7, -5, -4, -3, -3, -3, -1, 1, 2, 2, 4, 5, 8, 8, 10, 10, 10, 10, 10, 10, 11, 11, 12, 13, 15, 16, 17, 21, 26, 26, 31, 37},
	},
	{
		input:    []int{10, -2, -3, 0, -12, -16, -11, -19, 7, -19, 8, -3, 12, -9, -6, 15, -5, 1, -4, 17, 3, -2, -1, 3, 13},
		expected: []int{-19, -19, -16, -12, -11, -9, -6, -5, -4, -3, -3, -2, -2, -1, 0, 1, 3, 3, 7, 8, 10, 12, 13, 15, 17},
	},
	{
		input:    []int{20, -10, -20, -12, -8, 18, -9, -28, -8, 14, 12, 23, -3, -25, -27, -20, 23, 35, -32, 58, -50, -6, 32, 2, 14, 30, -55, -37, 15, 18, 18, 35, 35, -8, -11, 2, -10, 20, 37, -33, -28, 37, 33, -12, 65, 28, -22, -13, -7, -18, -40, 8, -51, 7, -38, 14, -37, 1, -32, 10, 6, -41, 1, -8, 19, 8, 54, 26, 19, 4, -62, 22, 13, -53, -18},
		expected: []int{-62, -55, -53, -51, -50, -41, -40, -38, -37, -37, -33, -32, -32, -28, -28, -27, -25, -22, -20, -20, -18, -18, -13, -12, -12, -11, -10, -10, -9, -8, -8, -8, -8, -7, -6, -3, 1, 1, 2, 2, 4, 6, 7, 8, 8, 10, 12, 13, 14, 14, 14, 15, 18, 18, 18, 19, 19, 20, 20, 22, 23, 23, 26, 28, 30, 32, 33, 35, 35, 35, 37, 37, 54, 58, 65},
	},
	{
		input:    []int{82, -77, 60, 80, 27, 11, -78, -68, 26, -32, -56, -92, 13, 6, -17, 87, -40, 2, 78, 9, 41, 15, -6, 57, -64, -17, 34, -13, 16, 28, 58, -23, 24, 4, 28, -69, -40, 52, 45, 10, 23, 78, -21, 8, -55, -23, -20, -52, -16, 52, -4, 71, 85, 78, -65, -85, 10, -54, 70, 28, 40, 60, -2, 27, 0, 30, -8, 7, -54, 67, -10, 26, 23, -72, 9, 6, 16, 71, -76, -12, -17, -24, 6, 44, -3, -5, -6, -2, -16, -29, 49, 26, 30, -95, 50, -35, 1, -42, -39, -63},
		expected: []int{-95, -92, -85, -78, -77, -76, -72, -69, -68, -65, -64, -63, -56, -55, -54, -54, -52, -42, -40, -40, -39, -35, -32, -29, -24, -23, -23, -21, -20, -17, -17, -17, -16, -16, -13, -12, -10, -8, -6, -6, -5, -4, -3, -2, -2, 0, 1, 2, 4, 6, 6, 6, 7, 8, 9, 9, 10, 10, 11, 13, 15, 16, 16, 23, 23, 24, 26, 26, 26, 27, 27, 28, 28, 28, 30, 30, 34, 40, 41, 44, 45, 49, 50, 52, 52, 57, 58, 60, 60, 67, 70, 71, 71, 78, 78, 78, 80, 82, 85, 87},
	},
	{
		input:    []int{332, 95, -141, -1, 49, -402, -401, -223, 103, 78, 23, -83, -101, 329, -7, -276, 169, 434, -65, -439, 127, -174, -246, -414, -8, -244, -95, -21, 110, 56, -51, -80, -34, 263, -156, 236, 38, -289, -38, -215, -141, -248, 17, -110, 248, -4, -154, -342, 249, -135, 71, 169, 3, 193, 119, -140, -341, -260, 234, -68, -440, -1, -102, 44, -41, 107, -156, 142, -277, -392, 126, 307, 245, 125, 47, 266, 283, 103, 345, -223, 251, 330, 201, 111, -97, 71, -350, 52, 217, -376, 158, -215, 69, 365, -130, 349, 215, -208, 300, 34, 12, 28, 25, 93, 14, -66, -154, -188, -323, 105, 129, 108, 58, 277, 256, 94, 112, -359, -286, 155, 12, 166, 10, 58, 34, -55, -352, 375, -18, -17, 361, 416, -272, -212, 297, 168, 14, 299, 390, -178, -279, -280, 199, 165, 257, -7, -361, -152, 409, 93, 269, 274, 164, -187, -314, -186, -109, -44, 106, 340, -71, -68, -204, -30, 220, -12, 222, 20, -55, -120, 230, -39, -150, 379, -274, -59, -237, -126, -19, -30, -110, -272, -23, 14, -212, -449, 60, 26, 224, 94, 324, 21, 228, 232, 169, 16, 99, 51, -112, -159, -342, -78, 156, 321, -309, 131, 244, -188, 444, -230, -238, -266, 39, 228, -117, -248, -158, 36, -337, 46, 306, 8, 133, -131, -77, -309, -203, 140, 7, 266, -266, -193, 404, -283, 434, -282, -274, 405, -147, -440, -308, -92, 356, 43, -121, 300, -250, -113, 293, -195, 36, 223, 42, -259, 49, -40, -149, -14, -349, -158, -305, 288, -252, 67, 226, -92, 207, 59, 122, 26, -83, -227, -21, 4, 217, -60, -114, 322, -255, -68, -129, -197, -258, -327, 210, 80, 71, 1, -85, 13, -72, 15, 164, -303, 186, -265, -338, -39, 163, 80, -16, 198, 213, -125, -4, -36, 379, -391, 135, -189, 16, 100, 334, 70, 300, -163, -300, 379, -234, 118, -118, 255, -50, -89, 126, -40, -146, -142, 214, -45, 84, 12, 394, 76, -124, 233, -259, 68, -234, 220, -72, -206, 98, -128, 217, 304, -137, 35, 91, -118, 280, 225, -263, -10, 388, -96, 104, 98, 104, 389, 12, -373, -337, -27, 11, 261, -12, -382, -1, 278, -244, -313, 30, 39, 369, -43, -392, 366, -276, -427, 122, 338, -208, 95, 37, -226, -142, 48, -339, 166, -440, -332, -33, -194, 270, -190, 231, -304, 364, 67, 264, 276, 107, 166, -261, -152, -214, 128, 110, -157, -175, 274, -234, -18, 101, -6, -45, -70, 287, -435, -66, 173, -348, 405, -38, 286, 110, -69, 25, -447, -74, 185, 41, 105, -107, -60, -154, -1, 30, -180, 26, -342, 403, -19, -147, -49, -272, -200, 42, 4, -6, 98, -40, -131, 0, 160, -110, 413, -296, -364, -211, -234, 53, 96, 239, 24, -224, 23, -356, -226, -430, -37, -89, 107, -79, -360, -199, 352, -19, -107, -178, -309, 51, -79, 44, 119, -102, -28, 82, -40, 88, 33, 341, -340, -19, -98, -398, 288, 169, -175},
		expected: []int{-449, -447, -440, -440, -440, -439, -435, -430, -427, -414, -402, -401, -398, -392, -392, -391, -382, -376, -373, -364, -361, -360, -359, -356, -352, -350, -349, -348, -342, -342, -342, -341, -340, -339, -338, -337, -337, -332, -327, -323, -314, -313, -309, -309, -309, -308, -305, -304, -303, -300, -296, -289, -286, -283, -282, -280, -279, -277, -276, -276, -274, -274, -272, -272, -272, -266, -266, -265, -263, -261, -260, -259, -259, -258, -255, -252, -250, -248, -248, -246, -244, -244, -238, -237, -234, -234, -234, -234, -230, -227, -226, -226, -224, -223, -223, -215, -215, -214, -212, -212, -211, -208, -208, -206, -204, -203, -200, -199, -197, -195, -194, -193, -190, -189, -188, -188, -187, -186, -180, -178, -178, -175, -175, -174, -163, -159, -158, -158, -157, -156, -156, -154, -154, -154, -152, -152, -150, -149, -147, -147, -146, -142, -142, -141, -141, -140, -137, -135, -131, -131, -130, -129, -128, -126, -125, -124, -121, -120, -118, -118, -117, -114, -113, -112, -110, -110, -110, -109, -107, -107, -102, -102, -101, -98, -97, -96, -95, -92, -92, -89, -89, -85, -83, -83, -80, -79, -79, -78, -77, -74, -72, -72, -71, -70, -69, -68, -68, -68, -66, -66, -65, -60, -60, -59, -55, -55, -51, -50, -49, -45, -45, -44, -43, -41, -40, -40, -40, -40, -39, -39, -38, -38, -37, -36, -34, -33, -30, -30, -28, -27, -23, -21, -21, -19, -19, -19, -19, -18, -18, -17, -16, -14, -12, -12, -10, -8, -7, -7, -6, -6, -4, -4, -1, -1, -1, -1, 0, 1, 3, 4, 4, 7, 8, 10, 11, 12, 12, 12, 12, 13, 14, 14, 14, 15, 16, 16, 17, 20, 21, 23, 23, 24, 25, 25, 26, 26, 26, 28, 30, 30, 33, 34, 34, 35, 36, 36, 37, 38, 39, 39, 41, 42, 42, 43, 44, 44, 46, 47, 48, 49, 49, 51, 51, 52, 53, 56, 58, 58, 59, 60, 67, 67, 68, 69, 70, 71, 71, 71, 76, 78, 80, 80, 82, 84, 88, 91, 93, 93, 94, 94, 95, 95, 96, 98, 98, 98, 99, 100, 101, 103, 103, 104, 104, 105, 105, 106, 107, 107, 107, 108, 110, 110, 110, 111, 112, 118, 119, 119, 122, 122, 125, 126, 126, 127, 128, 129, 131, 133, 135, 140, 142, 155, 156, 158, 160, 163, 164, 164, 165, 166, 166, 166, 168, 169, 169, 169, 169, 173, 185, 186, 193, 198, 199, 201, 207, 210, 213, 214, 215, 217, 217, 217, 220, 220, 222, 223, 224, 225, 226, 228, 228, 230, 231, 232, 233, 234, 236, 239, 244, 245, 248, 249, 251, 255, 256, 257, 261, 263, 264, 266, 266, 269, 270, 274, 274, 276, 277, 278, 280, 283, 286, 287, 288, 288, 293, 297, 299, 300, 300, 300, 304, 306, 307, 321, 322, 324, 329, 330, 332, 334, 338, 340, 341, 345, 349, 352, 356, 361, 364, 365, 366, 369, 375, 379, 379, 379, 388, 389, 390, 394, 403, 404, 405, 405, 409, 413, 416, 434, 434, 444},
	},
}
