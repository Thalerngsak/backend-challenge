package main

import (
	"fmt"
	"strings"
)

func findSmallestSumNumber(encoded string) string {
	left, right, equal := 1, 0, 1
	result := ""

	for _, r := range encoded {
		switch r {
		case 'L':
			result += fmt.Sprintf("%d%d", left, right)
			left++
		case 'R':
			result += fmt.Sprintf("%d%d", right, left)
			right++
		case '=':
			result += fmt.Sprintf("%d%d", equal, equal)
			equal++
		}
	}

	return result
}

func main() {
	var encoded string
	fmt.Print("Enter the encoded text data: ")
	fmt.Scanln(&encoded)

	encoded = strings.ToUpper(encoded)
	result := findSmallestSumNumber(encoded)
	fmt.Printf("Output: %s\n", result)
}
