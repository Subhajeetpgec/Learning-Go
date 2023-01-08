//Write a function that takes a slice of strings and returns a new slice with all strings that are
// palindromes (i.e., read the same backwards as forwards). 
//Input: ["racecar", "level", "hello"] 
//Output: ["racecar", "level"]


package main

import "fmt"

func main() {
	strs := []string{"racecar", "level", "hello"}
	palindromes := palindromeStrings(strs)
	fmt.Println(palindromes)
}

func palindromeStrings(strs []string) []string {
	// Create a slice to hold the palindrome strings
	palindromes := make([]string, 0)

	// Iterate over the input slice of strings
	for _, str := range strs {
		// Check if the string is a palindrome
		if isPalindrome(str) {
			// If it is a palindrome, append it to the slice
			palindromes = append(palindromes, str)
		}
	}

	// Return the slice of palindrome strings
	return palindromes
}

// Helper function to check if a string is a palindrome
func isPalindrome(s string) bool {
	// Convert the string to a slice of runes
	runes := []rune(s)

	// Check if the string is a palindrome by comparing the runes at the start and end
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

