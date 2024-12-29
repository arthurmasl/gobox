package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

func main() {
	// for n := range fibIterator(20) {
	// 	fmt.Println(n)
	// }

	for p := range removeLettersIterator("there is a missing letter") {
		fmt.Println(p)
	}
}

func removeLetters(phrase string) {
	letters := []rune{}

	for _, char := range phrase {
		letters = append(letters, char)
	}

	slices.Sort(letters)
	letters = slices.Compact(letters)

	for _, letter := range letters[1:] {
		newPhrase := strings.ReplaceAll(phrase, string(letter), "")
		phrase = newPhrase
	}
}

func removeLettersIterator(phrase string) iter.Seq[string] {
	return func(yield func(string) bool) {
		letters := []rune{}

		for _, char := range phrase {
			letters = append(letters, char)
		}

		slices.Sort(letters)
		letters = slices.Compact(letters)

		for _, letter := range letters[1:] {
			newPhrase := strings.ReplaceAll(phrase, string(letter), "")
			phrase = newPhrase
			if !yield(phrase) {
				return
			}
		}
	}
}

func rangeIterator(from, to int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range to - from + 1 {
			if !yield(i + from) {
				return
			}
		}
	}
}

func fibIterator(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			if !yield(a) {
				return
			}
		}
	}
}
