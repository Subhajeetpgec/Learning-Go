//Write a function that takes a string and returns a map with the frequency count of each character in the string. 
//Input: "hello" 
//Output: {"h": 1, "e": 1, "l": 2, "o": 1}

package main

import "fmt"

func main() {
	s := "hello"
	freq := characterFrequency(s)
	fmt.Println(freq) // Output: map[h:1 e:1 l:2 o:1]
}
func characterFrequency(s string) map[rune]int {
	// Create a map to hold the frequency count of each character
	freq := make(map[rune]int)

	// Iterate over the string and count the frequency of each character
	for _, r := range s {
		freq[r]++
	}

	return freq
}
