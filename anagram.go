//https://neetcode.io/problems/is-anagram
package main

import "fmt"

func isAnagram(s string, t string) bool {
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, ch := range s {
		count1[ch]++
	}
	for _, ch := range t {
		count2[ch]++
	}

	for key, val := range count1 {
		if count2[key] != val {
			return false
		}
	}
	return true

}

func main() {
	s := "hello"
	t := "holle"
	fmt.Println(isAnagram(s, t))
}