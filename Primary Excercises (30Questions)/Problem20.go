//Write a function that takes a slice of integers and returns a new slice with all the prime numbers. 
//Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] 
//Output: [2, 3, 5, 7]
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	primes := filterPrimes(nums)
	fmt.Println(primes) 
}
func filterPrimes(nums []int) []int {
	// Initialize a new slice to hold the filtered numbers
	filtered := make([]int, 0)
	// Iterate over the input slice and add prime numbers to the new slice
	for _, n := range nums {
		if isPrime(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

// Returns true if the given number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
