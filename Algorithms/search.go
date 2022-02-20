package main

import "fmt"

// BinarySearch works when the array is sorted
func BinarySearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		if x == arr[m] {
			return m
		} else if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}

func main() {
	items := []int{2, 1, 3, 4, 5, 6, 7, 8}
	fmt.Println(BinarySearch(items, 1))
}
