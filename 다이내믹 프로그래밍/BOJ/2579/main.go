package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/2579 - 계단 오르기

var memo = make(map[int]int)

func dp(stairs []int, n int) int {
	if value, ok := memo[n]; ok {
		return value
	}
	if n == 0 {
		return stairs[0]
	}
	if n == 1 {
		return stairs[0] + stairs[1]
	}
	if n == 2 {
		return max(stairs[0]+stairs[2], stairs[1]+stairs[2])
	}
	// 두 계단을 한꺼번에
	case1 := dp(stairs, n-2) + stairs[n]
	// 1칸씩 두 번 올라가는 경우
	case2 := dp(stairs, n-3) + stairs[n-1] + stairs[n]
	memo[n] = max(case1, case2)
	return memo[n]
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscan(reader, &N); err != nil {
		log.Fatal(err)
	}

	stairs := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &stairs[i]); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintln(writer, dp(stairs, N-1))
}
