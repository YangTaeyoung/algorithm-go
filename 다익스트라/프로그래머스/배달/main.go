package main

import (
	"container/heap"
	"fmt"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/12978 - 배달

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

func solution(N int, roads [][]int, k int) int {
	// 무향 그래프 초기화
	graph := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		graph[i] = make([]int, N+1)
		for j := 0; j < N+1; j++ {
			graph[i][j] = INF
		}
	}

	for _, road := range roads {
		if graph[road[0]][road[1]] > road[2] {
			graph[road[0]][road[1]] = road[2]
		}
		if graph[road[1]][road[0]] > road[2] {
			graph[road[1]][road[0]] = road[2]
		}
	}

	// 힙 및 구성요소 초기화
	pq := &EdgeHeap{
		{
			DST:  1,
			Cost: 0,
		},
	}
	heap.Init(pq)
	visited := make([]bool, N+1)
	costs := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		costs[i] = INF
	}
	costs[1] = 0

	// 다익스트라
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Edge)
		visited[current.DST] = true

		for next, cost := range graph[current.DST] {
			if visited[next] {
				continue
			}
			if c := current.Cost + cost; c < costs[next] {
				costs[next] = c
				heap.Push(pq, Edge{
					DST:  next,
					Cost: c,
				})
			}
		}
	}

	// 정답 계수
	answer := 0
	for _, cost := range costs {
		if k >= cost {
			answer++
		}
	}

	return answer
}

func main() {
	case1 := solution(5, [][]int{
		{1, 2, 1},
		{2, 3, 3},
		{5, 2, 2},
		{1, 4, 2},
		{5, 3, 1},
		{5, 4, 1},
	}, 3)
	fmt.Println(case1)

	case2 := solution(6, [][]int{
		{1, 2, 1},
		{1, 3, 2},
		{2, 3, 2},
		{3, 4, 3},
		{3, 5, 2},
		{3, 5, 3},
		{5, 6, 1},
	}, 4)
	fmt.Println(case2)
}
