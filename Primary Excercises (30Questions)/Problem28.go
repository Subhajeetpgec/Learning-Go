//Write a function that takes a slice of integers and returns the smallest number that is
//divisible by all elements in the slice.
//Input: [2, 3, 5]
// Output: 30

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 5}
	smallest := smallestMultiple(nums)
	fmt.Println(smallest)
}

func smallestMultiple(nums []int) int {
	// Sort the input slice in ascending order
	sort.Ints(nums)
	// Initialize the smallest multiple to the largest element in the slice
	smallest := nums[len(nums)-1]
	// Keep incrementing the smallest multiple until it is divisible by all elements in the slice
	for {
		found := true
		for _, num := range nums {
			if smallest % num != 0 {
				found = false
				break
			}
		}
		if found {
			break
		}
		smallest++
	}
	return smallest
}
