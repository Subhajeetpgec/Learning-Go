//Write a function that takes a slice of integers and returns a new slice with the elements in reverse order.
// Input: [1, 2, 3, 4, 5] 
//Output: [5, 4, 3, 2, 1]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	reversed := reverse(nums)
	fmt.Println(reversed) // Output: [5 4 3 2 1]
}
func reverse(nums []int) []int {
	// Initialize a new slice to hold the reversed elements
	reversed := make([]int, len(nums))
	// Iterate over the input slice and add the elements in reverse order to the new slice
	for i, n := range nums {
		reversed[len(nums)-1-i] = n
	}
	return reversed
}
