package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/43162 - 네트워크

func DFS(adj [][]int, visited *[]bool, node int) bool {
	if (*visited)[node] {
		return false
	}
	(*visited)[node] = true
	for _, n := range adj[node] {
		DFS(adj, visited, n)
	}

	return true
}

func solution(n int, computers [][]int) int {
	// 인접 리스트 초기화
	adj := make([][]int, n)
	for i, row := range computers {
		for j, connected := range row {
			if connected == 1 {
				adj[i] = append(adj[i], j)
			}
		}
	}

	answer := 0
	visited := make([]bool, n)
	for i := 0; i < len(computers); i++ {
		if DFS(adj, &visited, i) {
			answer++
		}
	}

	return answer
}

func main() {
	case1 := solution(3, [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	})
	fmt.Println(case1)
	case2 := solution(3, [][]int{
		{1, 1, 0}, {1, 1, 1}, {0, 1, 1},
	})
	fmt.Println(case2)
}
