//Write a function that takes a slice of integers and returns a new slice with all elements that are 
//divisible by a given number. 
//Input: ([1, 2, 3, 4, 5], 2) 
//Output: [2, 4]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	filtered := filterDivisible(nums, 2)
	fmt.Println(filtered)
}
func filterDivisible(nums []int, divisor int) []int {
	// Initialize a new slice to hold the filtered numbers
	filtered := make([]int, 0)
	// Iterate over the input slice and add numbers that are divisible by the given number to the new slice
	for _, n := range nums {
		if n%divisor == 0 {
			filtered = append(filtered, n)
		}
	}
	return filtered
}
