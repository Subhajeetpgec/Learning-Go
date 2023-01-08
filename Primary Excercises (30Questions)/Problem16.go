//Write a function that takes a slice of strings and a string, and returns the index of the 
//first occurrence of the string in the slice. 
//Input: (["cat", "dog", "elephant"], "dog") 
//Output: 1

package main

import "fmt"

func main() {
	slice := []string{"cat", "dog", "elephant"}
	index := findIndex(slice, "dog")
	fmt.Println(index)
}
func findIndex(slice []string, str string) int {
	for i, s := range slice {
		if s == str {
			return i
		}
	}
	return -1
}
