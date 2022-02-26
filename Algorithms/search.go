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

func FastLinearSearch(arr []int, x int) int {
	i := 0
	count := len(arr)
	arr = append(arr, x)
	for true {
		if arr[i] == x {
			if i < count {
				return i
			}
			return -1
		}
		i++
	}
	return -1
}

//InterpolationSearch works when the array is sorted
func InterpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for (arr[low] < x) && (x < arr[high]) {
		var mid = low + ((x-arr[low])*(high-low))/(arr[high]-arr[low])
		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	if arr[low] == x {
		return low
	}
	if arr[high] == x {
		return high
	}
	return -1
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(BinarySearch(items, 1))
	fmt.Println(FastLinearSearch(items, 2))
	fmt.Println(InterpolationSearch(items, 2))
}
