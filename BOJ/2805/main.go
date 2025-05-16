package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func cuttingTreeSum(trees []int, cut int) int {
	total := 0
	for _, tree := range trees {
		if tree-cut > 0 {
			total += tree - cut
		}
	}
	return total
}

func max(trees []int) int {
	m := 0
	for _, tree := range trees {
		if tree > m {
			m = tree
		}
	}

	return m
}

// https://www.acmicpc.net/problem/2805 - 나무자르기
func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, needs int
	if _, err := fmt.Fscan(reader, &N, &needs); err != nil {
		log.Fatal(err)
	}

	trees := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &trees[i]); err != nil {
			log.Fatal(err)
		}
	}

	minHeight := 0
	maxHeight := max(trees)
	middle := 0
	answer := 0
	for minHeight <= maxHeight {
		middle = (maxHeight + minHeight) / 2
		if sum := cuttingTreeSum(trees, middle); sum >= needs {
			answer = middle
			minHeight = middle + 1
		} else {
			maxHeight = middle - 1
		}
	}

	if _, err := fmt.Fprintln(writer, answer); err != nil {
		log.Fatal(err)
	}
}
