package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/11399 - ATM
func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	times := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &times[i]); err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(times)
	prefixSums := make([]int, N)
	total := 0
	for i, time := range times {
		if i == 0 {
			prefixSums[i] = time

		} else {
			prefixSums[i] = prefixSums[i-1] + time
		}
		total += prefixSums[i]
	}

	if _, err := fmt.Fprintln(writer, total); err != nil {
		log.Fatal(err)
	}

}
