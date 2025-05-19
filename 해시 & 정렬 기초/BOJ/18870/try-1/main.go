package main

// https://www.acmicpc.net/problem/18870 좌표 압축

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func uniq(s []int) map[int]bool {
	u := make(map[int]bool)
	for n := range s {
		u[s[n]] = true
	}
	return u
}

// findRank에서 좌표마다 u를 한번씩 더 도는 과정에서 O(N)시간이 초과하게 되어 시간초과로 실패함
func findRank(u map[int]bool, target int) int {
	cnt := 0
	for n := range u {
		if target > n {
			cnt++
		}
	}

	return cnt
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
	for i := 0; i < N; i++ {
		_, _ = fmt.Fprint(writer, findRank(u, coordinates[i]), " ")
	}

}
