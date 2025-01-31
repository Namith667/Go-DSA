package main

import (
	"fmt"
)

func main() {

	a := []int{13, 25, 67, 87, 54, 32}
	target := 25
	idx := binarySearch(a, target)
	if idx != -1 {
		fmt.Printf("found at %d", idx)
	} else {
		fmt.Println("Not Found")
	}
}

func binarySearch(a []int, target int) int {

	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2

		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			left = mid + 1
		} else if a[mid] > target {
			right = mid - 1
		}
	}
	return -1

}
