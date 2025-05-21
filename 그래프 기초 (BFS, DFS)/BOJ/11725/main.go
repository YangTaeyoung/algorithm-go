package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

// https://www.acmicpc.net/problem/11725 - 트리의 부모 찾기

func bfs(adj [][]int, visited []bool, node int) []int {
	queue := []int{node}
	parents := make([]int, len(adj))

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		visited[top] = true

		for _, n := range adj[top] {
			if !visited[n] && !slices.Contains(queue, n) {
				queue = append(queue, n)
				parents[n] = top
			}
		}
	}
	return parents
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatalln(err)
	}

	adj := make([][]int, N)
	for i := 0; i < N-1; i++ {
		var node, connected int

		if _, err := fmt.Fscanln(reader, &node, &connected); err != nil {
			log.Fatal(err)
		}
		adj[node-1] = append(adj[node-1], connected-1)
		adj[connected-1] = append(adj[connected-1], node-1)
	}

	parents := bfs(adj, make([]bool, N), 0)

	for i, parent := range parents {
		if i == 0 {
			continue
		}
		_, _ = fmt.Fprintln(writer, parent+1)
	}
}
