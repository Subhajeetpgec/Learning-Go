//Write a function that takes a string and returns a new string with all the words in reverse order. 
//Input: "This is a test" 
//Output: "test a is This"

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a test"
	reversed := reverseWords(s)
	fmt.Println(reversed) // Output: "test a is This"
}
func reverseWords(s string) string {
	// Split the string into an array of words
	words := strings.Split(s, " ")

	// Reverse the array of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed array of words into a single string
	return strings.Join(words, " ")
}
