package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/42577 - 전화번호 목록

func solution(phoneBook []string) bool {
	sort.Slice(phoneBook, func(i, j int) bool {
		return len(phoneBook[i]) < len(phoneBook[j])
	})

	for i, number := range phoneBook[:len(phoneBook)-1] {
		for _, nextNumber := range phoneBook[i+1:] {
			if strings.HasPrefix(nextNumber, number) {
				return false
			}
		}
	}

	return true
}

func main() {
	case1 := solution([]string{"119", "97674223", "1195524421"})
	fmt.Println(case1)
	case2 := solution([]string{"123", "456", "789"})
	fmt.Println(case2)
	case3 := solution([]string{"12", "123", "1235", "567", "88"})
	fmt.Println(case3)
}
