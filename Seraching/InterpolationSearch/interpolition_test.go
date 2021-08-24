package interpolition

import (
	"testing"
)

func TestInterpolitionSearch(t *testing.T) {
	t.Run("sampleDataset", func(t *testing.T) {
		for wantedNumber, dataSample := range sampleDataset {
			result := Search(dataSample, wantedNumber)
			if result < len(dataSample) && dataSample[result] == wantedNumber {
				t.Logf("Data Sample: %v\nNumber %d found in %d index.\n", dataSample, wantedNumber, result)
			} else {
				t.Fatalf("Data Sample: %v\nNumber %d can't found in Data Sample.\n", wantedNumber, dataSample)
			}
		}
	})
}

func BenchmarkInterpolitionSearch(b *testing.B) {
	dataSample := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	wantedNumber := 36

	for i := 0; i < b.N; i++ {
		result := Search(dataSample, wantedNumber)
		if !(result < len(dataSample) && dataSample[result] == wantedNumber) {
			b.Fatalf("Data Sample: %v\nNumber %d can't found in Data Sample.\n", wantedNumber, dataSample)
		}
	}
}

var sampleDataset = map[int][]int{
	45:  {1, 3, 6, 10, 15, 21, 28, 36, 45, 55},
	129: {1, 43, 52, 78, 110, 129, 160, 172, 188, 201},
	2:   {0, 2, 14, 33, 67, 79, 91, 112, 142, 149},
	902: {24, 31, 47, 69, 89, 103, 210, 578, 901, 902},
	12:  {12, 37, 62, 102, 151, 213, 285, 362, 455, 551},
}
