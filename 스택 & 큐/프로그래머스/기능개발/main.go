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

func solution(progresses []int, speeds []int) []int {
	restDays := make([]int, len(progresses))
	for i, progress := range progresses {
		restDays[i] = (100 - progress) / speeds[i]
		if (100-progress)%speeds[i] > 0 {
			restDays[i]++
		}
	}

	s := make(stack, 0)
	day := 0
	answer := make([]int, 0)
	for _, restDay := range restDays {
		if day != 0 && day < restDay {
			feature := 0
			for len(s) > 0 {
				s.pop()
				feature++
			}
			answer = append(answer, feature)
		}
		if len(s) == 0 {
			day = restDay
		}
		s.push(restDay)
	}
	if len(s) > 0 {
		answer = append(answer, len(s))
	}
	return answer
}

func main() {
	case1 := solution([]int{93, 30, 55}, []int{1, 30, 5})
	fmt.Println(case1)
	case2 := solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1})
	fmt.Println(case2)
}
