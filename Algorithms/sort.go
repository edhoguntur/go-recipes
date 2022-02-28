package main

import (
	"fmt"
	"math"
)

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

// Time Complexity 0(n+k)
// Space Complexity 0(k)

func CountingSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)

	var min = math.MaxInt32
	var max = math.MinInt32
	for _, x := range arr {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	var counts = make([]int, max-min+1)
	for _, x := range arr {
		counts[x-min]++
	}

	var total = 0
	for i := min; i <= max; i++ {
		var oldCount = counts[i-min]
		counts[i-min] = total
		total += oldCount
	}

	for _, x := range arr {
		items[counts[x-min]] = x
		counts[x-min]++
	}
	return items
}

func main() {
	items := []int{2, 6, 4, 5, 7, 1, 2}
	sortItems := BubbleSort(items)

	fmt.Println(sortItems)

	sortItems = CountingSort(items)

	fmt.Println(sortItems)
}
