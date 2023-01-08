//Write a function that takes a slice of integers and returns the index of the first occurrence of a given number. 
//Input: ([1, 2, 3, 2, 4, 5], 2) 
//Output: 1

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5}
	count := countNumber(nums, 2)
	fmt.Println(count) 
}

func countNumber(nums []int, num int) int {
	var count int
	for _, n := range nums {
		if n == num {
			count++
		}
	}
	return count
}
