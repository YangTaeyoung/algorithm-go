package main

// https://school.programmers.co.kr/learn/courses/30/lessons/42576

import (
	"fmt"
	"slices"
)

func solution(participants []string, completions []string) string {
	for _, participant := range participants {
		if !slices.Contains(completions, participant) {
			return participant
		} else {
			idx := slices.Index(completions, participant)
			completions = append(completions[:idx], completions[idx+1:]...)
		}
	}

	panic("No participant found")
}

func main() {
	// Test Case 1
	participants := []string{"leo", "kiki", "eden"}
	completions := []string{"eden", "kiki"}
	fmt.Println(solution(participants, completions))

	// Test Case 2
	participants = []string{"marina", "josipa", "nikola", "vinko", "filipa"}
	completions = []string{"josipa", "filipa", "marina", "nikola"}
	fmt.Println(solution(participants, completions))

	// Test Case 3
	participants = []string{"mislav", "stanko", "mislav", "ana"}
	completions = []string{"stanko", "ana", "mislav"}
	fmt.Println(solution(participants, completions))
}
