package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

// https://www.acmicpc.net/problem/1654 - 랜선 자르기

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var (
		K int
		N int
	)
	if _, err := fmt.Fscanln(reader, &K, &N); err != nil {
		log.Fatal(err)
	}

	cables := make([]int, K)
	for i := 0; i < K; i++ {
		if _, err := fmt.Fscanln(reader, &cables[i]); err != nil {
			log.Fatal(err)
		}
	}

	maxLength := slices.Max(cables) + 1
	minLength := 2

	answer := 0
	for minLength <= maxLength {
		total := 0
		mid := (minLength + maxLength) / 2
		for _, cable := range cables {
			total += cable / mid
		}
		if total >= N {
			minLength = mid + 1
			answer = mid
		}
		if total < N {
			maxLength = mid - 1
		}
	}

	if _, err := fmt.Fprintln(writer, answer); err != nil {
		log.Fatal(err)
	}
}
