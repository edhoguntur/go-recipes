package main

import "testing"

func BenchmarkBinarySearch(b *testing.B) {
	// Data Test
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
	for i := 0; i < b.N; i++ {
		BinarySearch(data, 2)
	}
}

func BenchmarkFastLinearSearch(b *testing.B) {
	// Data Test
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
	for i := 0; i < b.N; i++ {
		FastLinearSearch(data, 2)
	}
}
