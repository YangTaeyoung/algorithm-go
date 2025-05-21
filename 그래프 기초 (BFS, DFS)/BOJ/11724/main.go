package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/11724 - 연결 요소의 개수
func dfs(adj [][]int, visited []bool, node int) {
	visited[node] = true
	for _, n := range adj[node] {
		if !visited[n] {
			dfs(adj, visited, n)
		}
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()
	var N, M int
	if _, err := fmt.Fscanln(reader, &N, &M); err != nil {
		log.Fatal(err)
	}

	adj := make([][]int, N)
	for i := 0; i < M; i++ {
		var node, connected int
		if _, err := fmt.Fscanln(reader, &node, &connected); err != nil {
			log.Fatal(err)
		}
		adj[node-1] = append(adj[node-1], connected-1)
		adj[connected-1] = append(adj[connected-1], node-1)
	}

	visited := make([]bool, N)
	answer := 0
	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs(adj, visited, i)
			answer++
		}
	}

	if _, err := fmt.Fprintln(writer, answer); err != nil {
		log.Fatal(err)
	}
}
