//https://neetcode.io/problems/duplicate-integer

package main

import "fmt"

func checkDup(nums []int, s int) bool {

	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false

}

func main() {
	nums := []int{1, 2, 3, 4, 4}
	s := len(nums)
	fmt.Println(checkDup(nums, s))
}