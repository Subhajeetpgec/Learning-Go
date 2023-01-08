//Write a function that takes a slice of strings and returns a new slice
// with all strings that start with a given letter. 
//Input: (["cat", "dog", "elephant"], "e") 
//Output: ["elephant"]

package main

import "fmt"

func main() {
	slice := []string{"cat", "dog", "elephant"}
	filtered := filterByLetter(slice, "e")
	fmt.Println(filtered)
}


func filterByLetter(slice []string, letter string) []string {
	var result []string
	for _, s := range slice {
		if s[0] == letter[0] {
			result = append(result, s)
		}
	}
	return result
}
