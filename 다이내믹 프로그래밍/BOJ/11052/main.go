package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/11052 - 카드 구매하기

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}
	p := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		if _, err := fmt.Fscan(reader, &p[i]); err != nil {
			log.Fatal(err)
		}
	}

	dp := make(map[int]int)
	dp[1] = p[1]
	for i := 2; i < N+1; i++ {
		maximum := 0
		for j := 1; j <= i; j++ {
			if maximum < dp[i-j]+p[j] {
				maximum = dp[i-j] + p[j]
			}
		}
		dp[i] = maximum
	}

	if _, err := fmt.Fprintln(writer, dp[N]); err != nil {
		log.Fatal(err)
	}
}
