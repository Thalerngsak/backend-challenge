package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read JSON file
	jsonFile, err := os.Open("../files/hard.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var input [][]int
	json.Unmarshal(byteValue, &input)

	// Calculate the most valuable path
	mostValuablePath := findMostValuablePath(input)

	fmt.Printf("The most valuable path is %d\n", mostValuablePath)
}

func findMostValuablePath(input [][]int) int {
	n := len(input)
	dp := make([][]int, n)

	// Initialize the DP table
	for i := range dp {
		dp[i] = make([]int, len(input[i]))
	}

	// Set the first element
	dp[0][0] = input[0][0]

	// Calculate the path value for the rest of the elements
	for i := 1; i < n; i++ {
		for j := 0; j < len(input[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + input[i][j]
			} else if j == len(input[i])-1 {
				dp[i][j] = dp[i-1][j-1] + input[i][j]
			} else {
				dp[i][j] = max(dp[i-1][j-1], dp[i-1][j]) + input[i][j]
			}
		}
	}

	// Find the most valuable path in the last row
	mostValuablePath := dp[n-1][0]
	for i := 1; i < len(dp[n-1]); i++ {
		mostValuablePath = max(mostValuablePath, dp[n-1][i])
	}

	return mostValuablePath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
