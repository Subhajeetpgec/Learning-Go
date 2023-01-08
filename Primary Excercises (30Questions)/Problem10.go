//Write a function that takes a slice of strings and a string, and returns true if the 
//string is contained in the slice. 
//Input: (["cat", "dog", "elephant"], "dog") 
//Output: true

package main

import "fmt"

func main() {
	animal := []string{"cat", "dog", "elephant"}
	s := "dog"
	result := contains(animal, s)
	fmt.Println(result) 
}
func contains(strs []string, s string) bool {
	// Iterate over the slice and check if the input string is contained in the slice
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
