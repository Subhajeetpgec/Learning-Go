//Write a function that takes a slice of integers and returns a new slice with only the even numbers. 
//Input: [1, 2, 3, 4, 5] 
//Output: [2, 4]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	evens := evenNumbers(nums)
	fmt.Println(evens) 
}
func evenNumbers(nums []int) []int {
	// Create a new slice to hold the even numbers
	evens := make([]int, 0)

	// Iterate over the input slice and add the even numbers to the new slice
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}

	return evens
}
