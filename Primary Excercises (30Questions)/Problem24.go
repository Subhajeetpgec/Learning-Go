//Write a function that takes a slice of strings and returns a new slice with all strings that are at least a given length. 
//Input: (["cat", "dog", "elephant", "lion"], 5) 
//Output: ["elephant"]

package main

import "fmt"

func main() {
	strings := []string{"cat", "dog", "elephant", "lion"}
	filtered := filterLongStrings(strings, 5)
	fmt.Println(filtered) 
}
func filterLongStrings(strings []string, minLength int) []string {
	// Initialize a new slice to hold the filtered strings
	filtered := make([]string, 0)
	// Iterate over the input slice and add strings that are at least the given length to the new slice
	for _, s := range strings {
		if len(s) >= minLength {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
