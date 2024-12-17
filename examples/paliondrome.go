package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("kayak"))
}

func isPalindrome(word string) bool {
	n := len(word)
	for i := range n / 2 {
		if word[i] != word[n-i-1] {
			return false
		}
	}

	return true
}
