//Write a function that takes a slice of integers and returns true if the slice is sorted in descending order. 
//Input: [5, 4, 3, 2, 1] 
//Output: true

package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1}
	sorted := isSortedDescending(nums)
	fmt.Println(sorted) 
}
func isSortedDescending(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}
