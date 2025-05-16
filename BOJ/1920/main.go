package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/1920 - 수 찾기

func bs(arr []int, start int, end int, target int) int {
	middle := (start + end) / 2
	if start == end {
		return 0
	}
	if arr[middle] == target {
		return 1
	} else if arr[middle] > target { // 타겟이 더 작음
		return bs(arr, start, middle, target)
	} else { // 타겟이 더 큼
		return bs(arr, middle+1, end, target)
	}
}

func main() {
	var (
		writer = bufio.NewWriter(os.Stdout)
		reader = bufio.NewReader(os.Stdin)
	)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscan(reader, &N); err != nil {
		log.Fatal(err)
	}

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			log.Fatal(err)
		}
	}

	var M int
	if _, err := fmt.Fscan(reader, &M); err != nil {
		log.Fatal(err)
	}

	targets := make([]int, M)
	for i := 0; i < M; i++ {
		if _, err := fmt.Fscan(reader, &targets[i]); err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(arr)

	for _, target := range targets {
		if _, err := fmt.Fprintln(writer, bs(arr, 0, len(arr), target)); err != nil {
			log.Fatal(err)
		}
	}

}
