package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/11047 - 동전 0
func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, K int
	if _, err := fmt.Fscan(reader, &N, &K); err != nil {
		log.Fatal(err)
	}

	coins := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &coins[i]); err != nil {
			log.Fatal(err)
		}
	}

	result := 0
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	for _, coin := range coins {
		result += K / coin
		K %= coin
	}

	fmt.Fprint(writer, result)
}
