package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type Edge struct {
	DST    int
	Weight int
}

type Edges []Edge

func (h *Edges) Len() int {
	return len(*h)
}

func (h *Edges) Less(i, j int) bool {
	return (*h)[i].Weight < (*h)[j].Weight
}

func (h *Edges) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Edges) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *Edges) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

const INF = 1e9

// https://www.acmicpc.net/problem/1753 - 최단경로
func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var V, E int
	if _, err := fmt.Fscanln(reader, &V, &E); err != nil {
		log.Fatal(err)
	}

	var K int
	if _, err := fmt.Fscanln(reader, &K); err != nil {
		log.Fatal(err)
	}

	graph := make([]Edges, V+1)
	for i := 0; i < E; i++ {
		var u, v, w int
		if _, err := fmt.Fscanln(reader, &u, &v, &w); err != nil {
			log.Fatal(err)
		}
		graph[u] = append(graph[u], Edge{DST: v, Weight: w})
	}

	visited := make([]bool, V+1)

	pq := &Edges{}
	heap.Init(pq)
	distances := make([]int, V+1)
	for node := 1; node <= V; node++ {
		if node == K {
			distances[node] = 0
		} else {
			distances[node] = INF
		}
	}

	pq.Push(Edge{DST: K, Weight: 0})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Edge)
		visited[current.DST] = true
		for _, edge := range graph[current.DST] {
			if visited[edge.DST] {
				continue
			}
			if distance := edge.Weight + current.Weight; distance < distances[edge.DST] {
				distances[edge.DST] = distance
				heap.Push(pq, Edge{DST: edge.DST, Weight: distance})
			}
		}
	}

	for i := 1; i <= V; i++ {
		if distances[i] == INF {
			if _, err := fmt.Fprintln(writer, "INF"); err != nil {
				log.Fatal(err)
			}
		} else {
			if _, err := fmt.Fprintln(writer, distances[i]); err != nil {
				log.Fatal(err)
			}
		}
	}
}
