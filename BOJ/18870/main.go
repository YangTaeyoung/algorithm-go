package main

// https://www.acmicpc.net/problem/18870 좌표 압축

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func uniq(s []int) []int {
	u := make(map[int]bool)
	for n := range s {
		u[s[n]] = true
	}

	arr := make([]int, 0, len(u))
	for k := range u {
		arr = append(arr, k)
	}

	return arr
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
	coordinates := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &coordinates[i]); err != nil {
			log.Fatal(err)
		}
	}

	u := uniq(coordinates)
	sort.Ints(u)

	rankMap := make(map[int]int, len(u))
	for rank, value := range u {
		rankMap[value] = rank
	}

	for i := 0; i < N; i++ {
		_, _ = fmt.Fprint(writer, rankMap[coordinates[i]], " ")
	}

}
