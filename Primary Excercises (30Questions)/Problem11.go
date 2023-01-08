//Write a function that takes a slice of integers and returns the second-largest number in the slice. 
//Input: [1, 2, 3, 4, 5] 
//Output: 4

 package main

 import "fmt"
 
 func main() {
	 nums := []int{1, 2, 3, 4, 5}
	 secondMax := secondMax(nums)
	 fmt.Println(secondMax)
 }
 func secondMax(nums []int) int {
	// Initialize variables to hold the maximum and second-maximum numbers
	max, secondMax := nums[0], nums[0]
	// Iterate over the slice and update the maximum and second-maximum numbers
	for _, n := range nums {
		if n > max {
			secondMax = max
			max = n
		} else if n > secondMax {
			secondMax = n
		}
	}
	return secondMax
}
