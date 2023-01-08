//Write a function that takes a slice of strings and returns a new slice with all strings 
//that have more than 5 characters. 
//Input: ["cat", "dog", "elephant", "lion"] 
//Output: ["elephant"]

package main

import "fmt"

func main() {
	animal := []string{"cat", "dog", "elephant", "lion"}
	long := longStrings(animal)
	fmt.Println(long) 
}
func longStrings(strs []string) []string {
	
	long := make([]string, 0) // Creating a new slice to hold the long strings
	// Iterate over the input slice and add the long strings to the new slice
	for _, s := range strs {
		if len(s) > 5 {
			long = append(long, s)
		}
	}
	return long
}
