package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/64061 - 크레인 인형뽑기

func solution(board [][]int, moves []int) int {
	var (
		stack  []int
		answer int
	)

	for _, move := range moves {
		for i := 0; i < len(board); i++ {
			if picked := board[i][move-1]; picked != 0 {
				if len(stack) > 0 && stack[len(stack)-1] == picked {
					answer += 2
					stack = stack[:len(stack)-1]
					break
				} else {
					board[i][move-1] = 0
					stack = append(stack, picked)
					break
				}
			}
		}
	}

	return answer
}

func main() {
	answer := solution(
		[][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 3},
			{0, 2, 5, 0, 1},
			{4, 2, 4, 4, 2},
			{3, 5, 1, 3, 1},
		},
		[]int{1, 5, 3, 5, 1, 2, 1, 4},
	)

	fmt.Println(answer)
}
