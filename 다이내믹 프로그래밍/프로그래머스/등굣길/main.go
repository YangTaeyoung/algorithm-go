package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/42898 - 등굣길

type Puddles [][]int

func (p Puddles) Contains(x, y int) bool {
	for _, puddle := range p {
		if puddle[0] == x && puddle[1] == y {
			return true
		}
	}
	return false
}

func solution(m int, n int, puddles Puddles) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[1][1] = 1

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if i == 1 && j == 1 || puddles.Contains(i, j) {
				continue
			}

			if i == 1 {
				dp[i][j] = dp[i][j-1]
			} else if j == 1 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(solution(4, 3, Puddles{
		{2, 2},
	}))
}
