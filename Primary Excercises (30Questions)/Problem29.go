//Write a function that takes a slice of integers and returns a new slice with all elements that are palindromes (i.e., read the same forwards as backwards). 
//Input: [1, 11, 121, 12321] 
//Output: [11, 121, 12321]


package main

import "fmt"

func main() {
	slice := []int{1, 11, 121, 12321}
	filtered := filterByPalindrome(slice)
	fmt.Println(filtered)
}

func filterByPalindrome(slice []int) []int {
	var result []int
	for _, n := range slice {
		if isPalindrome(n) {
			result = append(result, n)
		}
	}
	return result
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return s == string(runes)
}
