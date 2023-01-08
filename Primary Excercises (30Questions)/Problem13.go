//Write a function that takes a slice of integers and a number and returns the number of 
//times the number appears in the slice. 
//Input: ([1, 2, 3, 2, 4, 5], 2)
// Output: 2

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5}
	n := 2
	count := count(nums, n)
	fmt.Println(count) 
}
func count(nums []int, n int) int {
	// Initialize a variable to hold the count
	count := 0
	// Iterate over the slice and increment the count for each occurrence of the given number
	for _, num := range nums {
		if num == n {
			count++
		}
	}
	return count
}
