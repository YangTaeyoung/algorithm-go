package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/9251 - LCS

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var str1, str2 string
	if _, err := fmt.Fscanln(reader, &str1); err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fscanln(reader, &str2); err != nil {
		log.Fatal(err)
	}

	dp := make([][]int, len(str1)+1)
	dp[0] = make([]int, len(str2)+1)
	for i := 1; i <= len(str1); i++ {
		dp[i] = make([]int, len(str2)+1)
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Fprintln(writer, dp[len(str1)][len(str2)])
}
