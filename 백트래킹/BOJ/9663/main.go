package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/9663 - N-Queen

var (
	answer = 0
	N      int
)

type Location struct {
	x, y int
}

func canPlaced(location Location, queens []Location) bool {
	for _, queen := range queens {
		if queen.x == location.x || queen.y == location.y {
			return false
		}

		if slope := float64(location.y-queen.y) / float64(location.x-queen.x); slope == -1 || slope == 1 {
			return false
		}
	}

	return true
}

func dfs(queens []Location) {
	y := len(queens) // 현재 놓을 퀸의 y좌표 = 몇 번째 행인지
	if y == N {
		answer++
		return
	}

	for x := 0; x < N; x++ {
		location := Location{x, y}
		if canPlaced(location, queens) {
			dfs(append(queens, location))
		}
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	dfs(make([]Location, 0, N))

	fmt.Println(answer)

}
