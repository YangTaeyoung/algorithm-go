package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/64062 - 징검다리 건너기

func check(stones []int, minus int, k int) bool {
	count := 0
	for _, stone := range stones {
		if stone-minus <= 0 {
			count++
			continue
		}
		if count >= k {
			return true
		}
		count = 0
	}

	return count >= k
}

func max(arr []int) int {
	m := 0
	for _, n := range arr {
		if m < n {
			m = n
		}
	}
	return m
}

func solution(stones []int, k int) int {
	maxStone := max(stones)
	minStone := 0
	answer := 0
	for minStone <= maxStone {
		mid := (minStone + maxStone) / 2
		if check(stones, mid, k) {
			answer = mid
			maxStone = mid - 1
		} else {
			minStone = mid + 1
		}
	}

	return answer
}

func main() {
	case1 := solution([]int{2, 4, 5, 3, 2, 1, 4, 2, 5, 1}, 3)
	fmt.Println(case1)
	case2 := solution([]int{2, 3, 2, 2}, 3)
	fmt.Println(case2)
}
