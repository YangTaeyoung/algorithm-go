package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/2217 - 로프

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	lopes := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscanln(reader, &lopes[i]); err != nil {
			log.Fatal(err)
		}
	}

	sort.Slice(lopes, func(i, j int) bool {
		return lopes[i] > lopes[j]
	})

	maxWeight := 0
	for i, lope := range lopes {
		if weight := (i + 1) * lope; weight > maxWeight {
			maxWeight = weight
		}

	}

	if _, err := fmt.Fprintln(writer, maxWeight); err != nil {
		log.Fatal(err)
	}
}
