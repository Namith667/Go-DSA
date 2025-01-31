package main

import "fmt"

func main() {

	a := []int{22, 33, 33, 44, 44, 66}
	target := 33
	firstRes := findFirst(a, target)
	lastRes := findLast(a, target)
	fmt.Printf("\nfirst occurance of %d at index: ", firstRes)
	fmt.Printf("\nlast occurance of %d at index: ", lastRes)
}

func findFirst(a []int, target int) int {
	left, right := 0, len(a)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			result = mid
			right = mid - 1
		} else if a[mid] < target {
			left = mid + 1
		} else if a[mid] > target {
			right = mid - 1
		}
	}
	return result
}
func findLast(a []int, target int) int {
	left, right := 0, len(a)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			result = mid
			left = mid + 1
		} else if a[mid] < target {
			left = mid + 1
		} else if a[mid] > target {
			right = mid - 1
		}
	}
	return result
}
