//Write a function that takes a slice of integers and returns the largest number in the slice. 
//Input: [1, 2, 3, 4, 5] 
//Output: 5

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	maximum := max(nums)
	fmt.Println(maximum) 
}
func max(nums []int) int {
	max := nums[0] //storing the max. value
	// Iterate over the slice and update the maximum number if a larger number is found
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
