//Write a function that takes a string and returns a new string with all vowels removed. 
//Input: "hello" 
//Output: "hll"

package main

import "fmt"

func main() {
	s := "hello"
	modified := removeVowels(s)
	fmt.Println(modified) 
}
func removeVowels(s string) string {
	
	modified := "" //new string to hold the modified string
	// Iterate over the input string and add only the consonants to the new string
	for _, r := range s {
		if r != 'a' && r != 'e' && r != 'i' && r != 'o' && r != 'u' {
			modified += string(r)
		}
	}
	return modified
}

