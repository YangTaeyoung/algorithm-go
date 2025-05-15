package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/42578 - 의상

func solution(clothes [][]string) int {
	clothesMap := make(map[string]int)
	for _, c := range clothes {
		clothesMap[c[1]]++
	}

	answer := 1
	for _, count := range clothesMap {
		answer *= count + 1
	}
	return answer - 1
}

func main() {
	clothes := [][]string{
		{"yellow_hat", "headgear"},
		{"blue_sunglasses", "eyewear"},
		{"green_turban", "headgear"},
	}
	fmt.Println(solution(clothes))
}
