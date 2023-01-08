//Write a function that takes a slice of integers and returns the average of all the elements in the slice.
// Input: [1, 2, 3, 4, 5] 
//Output: 3

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	avg := average(nums)
	fmt.Println(avg) 
}
func average(nums []int) float64 {
	// Initialize a variable to hold the sum of the elements
	var sum int
	// Iterate over the slice and add each element to the sum
	for _, n := range nums {
		sum += n
	}
	// Calculate the average by dividing the sum by the number of elements
	return float64(sum) / float64(len(nums))
}
