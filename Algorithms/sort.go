package main

import (
	"fmt"
	"math"
)

// Time Complexity O(n^2)
// Space Complexity O(1)

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

// Time Complexity O(n+k)
// Space Complexity O(k)

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

// Time Complexity O(n log(n))
// Space Complexity O(n)

func MergeSort(arr []int) []int {
	items := make([]int, len(arr))
	copy(items, arr)
	DoMergeSort(items)
	return items
}

func DoMergeSort(items []int) {
	length := len(items)
	if length == 1 {
		return
	}
	var lLeft = length / 2
	var left = make([]int, lLeft)
	copy(left, items[:lLeft])
	var lRight = length - lLeft
	var right = make([]int, lRight)
	copy(right, items[lLeft:])

	DoMergeSort(left)
	DoMergeSort(right)

	Merge(left, right, items)
}

func Merge(left []int, right []int, result []int) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	var length = len(left) - l
	copy(result[i:i+length], left[l:])
	i = i + length
	length = len(right) - r
	copy(result[i:i+length], right[r:])
}

// Time Complexity from O(n log(n)) to O(n^2)
// Space Complexity O(log(n))

func DoSort(items []int, first int, last int) {
	if first >= last {
		return
	}
	i := first
	j := last
	x := items[(first+last)/2]

	for i < j {
		for items[i] < x {
			i++
		}
		for items[j] > x {
			j--
		}
		if i <= j {
			var tmp = items[i]
			items[i] = items[j]
			items[j] = tmp
			i++
			j--
		}
	}
	DoSort(items, first, j)
	DoSort(items, i, last)
}

func QuickSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	DoSort(items, 0, length-1)
	return items
}

func main() {
	items := []int{2, 6, 4, 5, 7, 1, 2}
	sortItems := BubbleSort(items)

	fmt.Println(sortItems)

	sortItems = CountingSort(items)

	fmt.Println(sortItems)

	sortItems = MergeSort(items)
	fmt.Println(sortItems)

	sortItems = QuickSort(items)
	fmt.Println(sortItems)
}
