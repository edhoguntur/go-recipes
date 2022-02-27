package main

import "fmt"

// Time Complexity 0(n^2)
// Space Complexity 0(1)

func BubbleSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if items[j] < items[i] {
				var tmp = items[j]
				items[j] = items[i]
				items[i] = tmp
			}
		}
	}
	return items
}

func main() {
	items := []int{2, 6, 4, 5, 7, 1, 2}
	sortItems := BubbleSort(items)

	fmt.Println(sortItems)
}
