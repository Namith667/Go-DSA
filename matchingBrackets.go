// Check Balanced Brackets
// Write a golang program to check whether the brackets are balanced
// for the given string or not. Eg. input [{}] then it should return
// true because every open bracket has a close bracket and if the
// input [[}] then it will return false because the opening
// bracket doesn’t have a closed bracket.
package main

import "fmt"

func main() {

	str := "[()}"
	fmt.Println(check(str))

}

func check(str string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range str {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
