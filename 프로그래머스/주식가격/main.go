package main

import "fmt"

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s *stack) top() int {
	if len(*s) == 0 {
		return -1
	}

	return (*s)[len(*s)-1]
}

func solution(tickers []int) []int {
	var s stack
	answer := make([]int, len(tickers))

	for i, ticker := range tickers {
		for len(s) != 0 && tickers[s.top()] > ticker {
			idx := s.pop()
			answer[idx] = i - idx
		}
		s.push(i)
	}

	for len(s) > 0 {
		idx := s.pop()
		answer[idx] = len(tickers) - 1 - idx
	}

	return answer
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 2, 3}))

}
