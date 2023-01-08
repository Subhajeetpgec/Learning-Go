//Write a function that takes a slice of integers and returns a new slice with the elements in alternating order. 
//Input: [1, 2, 3, 4, 5]
// Output: [1, 3, 5, 2, 4]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	alternating := alternate(nums)
	fmt.Println(alternating) 
}

func alternate(nums []int) []int {
	// Initialize a new slice to hold the alternating elements
	alternating := make([]int, 0)
	// Iterate over the input slice and add the elements in alternating order to the new slice
	for i, n := range nums {
		if i%2 == 0 {
			alternating = append(alternating, n)
		}
	}
	for i, n := range nums {
		if i%2 != 0 {
			alternating = append(alternating, n)
		}
	}
	return alternating
}
