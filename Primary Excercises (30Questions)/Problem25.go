//Write a function that takes a slice of strings and returns a new slice with all strings that 
//contain a given letter. 
//Input: (["cat", "dog", "elephant", "lion"], "e") 
//Output: ["elephant"]

package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"cat", "dog", "elephant", "lion"}
	stringsWithE := stringsContainingLetter(strs, "e")
	fmt.Println(stringsWithE) // Output: ["elephant"]
}

func stringsContainingLetter(strs []string, letter string) []string {
	// Create a slice to hold the strings that contain the given letter
	stringsWithLetter := make([]string, 0)

	// Iterate over the input slice of strings
	for _, str := range strs {
		// Check if the string contains the given letter
		if strings.Contains(str, letter) {
			// If it does, append it to the slice
			stringsWithLetter = append(stringsWithLetter, str)
		}
	}

	// Return the slice of strings that contain the given letter
	return stringsWithLetter
}
