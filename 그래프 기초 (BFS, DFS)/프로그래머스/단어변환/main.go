package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/43163 - 단어 변환

type Queue []string

func (q *Queue) Push(v string) {
	for _, word := range *q {
		if word == v {
			return
		}
	}

	*q = append(*q, v)
}

func isConnected(str1 string, str2 string) bool {
	cnt := 0
	for i := range str1 {
		if str1[i] == str2[i] {
			cnt++
		}
	}
	if cnt >= len(str1)-1 {
		return true
	}
	return false
}

func solution(begin string, target string, words []string) int {
	words = append(words, begin)
	adj := make(map[string][]string)

	for i, word1 := range words {
		for _, word2 := range words[i+1:] {
			if isConnected(word1, word2) && word1 != word2 {
				adj[word1] = append(adj[word1], word2)
				adj[word2] = append(adj[word2], word1)
			}
		}
	}

	queue := Queue{begin}

	count := make(map[string]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range adj[current] {
			if count[next] == 0 {
				queue.Push(next)
				count[next] = count[current] + 1
			}
		}
	}

	return count[target]
}

func main() {
	fmt.Println(solution("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(solution("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
