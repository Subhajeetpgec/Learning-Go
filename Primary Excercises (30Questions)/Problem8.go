//Write a function that takes a slice of integers and returns a new slice with all duplicates removed. 
//Input: [1, 2, 3, 2, 4, 5, 5] 
//Output: [1, 3, 4]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 5}
	filtered := removeDuplicates(nums)
	fmt.Println(filtered) 
}

func removeDuplicates(nums []int) []int {
	// Create a map to hold the unique numbers
	unique := make(map[int]bool)
	// Iterate over the input slice and add the unique numbers to the map
	for _, n := range nums {
		unique[n] = true
	}
	
	filtered := make([]int, 0) // Create a new slice to hold the unique numbers
	// Iterate over the map and add the keys (which are the unique numbers) to the new slice
	for k := range unique {
		filtered = append(filtered, k)
	}
	return filtered
}
