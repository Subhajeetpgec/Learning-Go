//Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order. 
//Input: [1, 2, 3, 4, 5] 
//Output: true

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sorted := isSorted(nums)
	fmt.Println(sorted) // Output: true
}
func isSorted(nums []int) bool {
	// Iterate over the slice and check if the current element is greater than or equal to the previous element
	for i := 1; i < len(nums); i++ { //i=1, 2 .....
		if nums[i] < nums[i-1] { //2<1, 3<2 ....
			return false
		}
	}
	return true
}
