package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/1916 - 최소비용 구하기

type Edge struct {
	DST  int
	Cost int
}

type EdgeHeap []Edge

func (e *EdgeHeap) Less(i, j int) bool {
	return (*e)[i].Cost < (*e)[j].Cost
}

func (e *EdgeHeap) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func (e *EdgeHeap) Len() int {
	return len(*e)
}

func (e *EdgeHeap) Push(x interface{}) {
	*e = append(*e, x.(Edge))
}

func (e *EdgeHeap) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

const INF = 1e9

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fscanln(reader, &M); err != nil {
		log.Fatal(err)
	}

	// 그래프 초기화
	graph := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		graph[i] = make([]int, N+1)
		for j := 0; j < N+1; j++ {
			graph[i][j] = INF
		}
	}

	for i := 0; i < M; i++ {
		var u, v, w int
		if _, err := fmt.Fscanln(reader, &u, &v, &w); err != nil {
			log.Fatal(err)
		}
		if w < graph[u][v] {
			graph[u][v] = w
		}
	}

	var src, dst int
	if _, err := fmt.Fscanln(reader, &src, &dst); err != nil {
		log.Fatal(err)
	}
	costs := make([]int, N+1)
	visited := make([]bool, N+1)
	for i := 0; i < N+1; i++ {
		costs[i] = INF
	}
	costs[src] = 0

	h := &EdgeHeap{
		{
			DST:  src,
			Cost: 0,
		},
	}
	heap.Init(h)

	for h.Len() > 0 {
		current := heap.Pop(h).(Edge)
		visited[current.DST] = true

		for next, cost := range graph[current.DST] {
			if visited[next] {
				continue
			}
			if c := current.Cost + cost; c < costs[next] {
				costs[next] = c
				heap.Push(h, Edge{
					DST:  next,
					Cost: c,
				})
			}
		}
	}

	if _, err := fmt.Fprintln(writer, costs[dst]); err != nil {
		log.Fatal(err)
	}
}
