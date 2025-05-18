package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/1260 - DFS와 BFS

func DFS(adj [][]int, visited []bool, start int) {
	visited[start] = true
	fmt.Printf("%d ", start+1)
	for _, node := range adj[start] {
		if !visited[node] {
			DFS(adj, visited, node)
		}
	}
}

func contains(q []int, e int) bool {
	for _, a := range q {
		if a == e {
			return true
		}
	}
	return false
}

func BFS(adj [][]int, visited []bool, start int) {
	q := []int{start}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		visited[current] = true
		fmt.Printf("%d ", current+1)
		for _, node := range adj[current] {
			if !visited[node] && !contains(q, node) {
				q = append(q, node)
			}
		}
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M, V int
	if _, err := fmt.Fscan(reader, &N, &M, &V); err != nil {
		log.Fatal(err)
	}
	adjArr := make([][]int, N)

	for i := 0; i < M; i++ {
		var node, adj int
		if _, err := fmt.Fscan(reader, &node, &adj); err != nil {
			log.Fatal(err)
		}
		adjArr[node-1] = append(adjArr[node-1], adj-1)
		adjArr[adj-1] = append(adjArr[adj-1], node-1)
	}
	for i := 0; i < N; i++ {
		sort.Ints(adjArr[i])
	}

	DFS(adjArr, make([]bool, N), V-1)
	fmt.Println()
	BFS(adjArr, make([]bool, N), V-1)
}
