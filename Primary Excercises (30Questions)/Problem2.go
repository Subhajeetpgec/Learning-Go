//Write a function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order.
// Input: ["zebra", "monkey", "aardvark"]
//Output: ["aardvark", "monkey", "zebra"]

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"zebra", "monkey", "aardvark"}
	sorted := sortStrings(strs)
	fmt.Println(sorted) 
}
func sortStrings(strs []string) []string {
	// Creating a new slice to hold the sorted strings
	
	sorted := make([]string, len(strs))
	// Copy the input slice to the new slice
	copy(sorted, strs)
	// Sorting the new slice
	sort.Strings(sorted)
	return sorted
}
