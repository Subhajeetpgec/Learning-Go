//Write a function that takes a slice of strings and returns a new slice 
//with all strings that are palindromes (i.e., read the same backwards as forwards). 
//Input: ["racecar", "level", "hello"] 
//Output: ["racecar", "level"]



func filterByPalindrome(slice []string) []string {
	var result []string
	for _, s := range slice {
		if s == reverse(s) {
			result = append(result, s)
		}
	}
	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
