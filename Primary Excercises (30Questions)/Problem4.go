//Write a function that takes a slice of integers and returns the sum of all the elements in the slice. 
//Input: [1, 2, 3, 4, 5] 
//Output: 15

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums)
	fmt.Println(total) 
}
func sum(nums []int) int {
	var sum int
	// Iterate over the slice and add each element to the total
	for _, n := range nums {
		sum += n
	}
	return sum
}
