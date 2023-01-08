//Write a function that takes a slice of integers and returns true if the slice contains a given number.
// Input: ([1, 2, 3, 4, 5], 3)
// Output: true
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	contains := contains(nums, 3)
	fmt.Println(contains)
}

func contains(nums []int, num int) bool {
	// Iterate over the input slice and return true if the given number is found
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	// Return false if the given number is not found
	return false
}
