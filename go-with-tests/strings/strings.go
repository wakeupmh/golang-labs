package strings

import "strings"


func Compare(a, b string) int {
	return strings.Compare(a, b) 
}

func ContainsWord(a, b string) bool {
	return strings.Contains(a, b)
}