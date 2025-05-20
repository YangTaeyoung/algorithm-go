package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

// https://www.acmicpc.net/problem/2606 - 바이러스
type Queue struct {
	top   int
	queue []int
}

func (q *Queue) Len() int {
	return len(q.queue) - q.top
}
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	q.top++
	return q.queue[q.top-1]
}

func (q *Queue) Push(v int) {
	if !slices.Contains(q.queue, v) {
		q.queue = append(q.queue, v)
	}
}

func bfs(adj [][]int, visited []bool, node int) int {
	var q Queue
	q.Push(node)
	result := 0
	for !q.IsEmpty() {
		current := q.Pop()
		visited[current] = true
		result++
		for _, n := range adj[current] {
			if visited[n] {
				continue
			}
			q.Push(n)
		}
	}

	return result
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var (
		computerCount int
		edgeCount     int
	)
	if _, err := fmt.Fscanln(reader, &computerCount); err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fscanln(reader, &edgeCount); err != nil {
		log.Fatal(err)
	}

	adj := make([][]int, computerCount)
	for i := 0; i < edgeCount; i++ {
		var node, adjacent int
		if _, err := fmt.Fscanln(reader, &node, &adjacent); err != nil {
			log.Fatal(err)
		}
		adj[node-1] = append(adj[node-1], adjacent-1)
		adj[adjacent-1] = append(adj[adjacent-1], node-1)
	}

	visited := make([]bool, computerCount)
	result := bfs(adj, visited, 0)

	if _, err := fmt.Fprintln(writer, result-1); err != nil {
		log.Fatal(err)
	}
}
