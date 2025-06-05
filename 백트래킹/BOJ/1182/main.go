package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/1182 - 부분 수열의 합

var (
	arr    []int
	answer int
	N      int
	S      int
)

func dfs(cur int, sum int) {
	if cur == N {
		if sum == S {
			answer++
		}
		return
	}

	dfs(cur+1, sum+arr[cur])
	dfs(cur+1, sum)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	if _, err := fmt.Fscanln(reader, &N, &S); err != nil {
		log.Fatal(err)
	}

	arr = make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			log.Fatal(err)
		}
	}
	if S == 0 {
		answer--
	}
	dfs(0, 0)

	fmt.Println(answer)
}
