package main

import (
	"fmt"
	"sort"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/42748 - K번째 수
func solution(array []int, commands [][]int) []int {
	answer := make([]int, len(commands))

	for i, command := range commands {
		s := append([]int{}, array[command[0]-1:command[1]]...)
		sort.Ints(s)
		answer[i] = s[command[2]-1]
	}

	return answer
}

func main() {
	fmt.Println(
		solution(
			[]int{1, 5, 2, 6, 3, 7, 4},
			[][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}},
		),
	)
}
