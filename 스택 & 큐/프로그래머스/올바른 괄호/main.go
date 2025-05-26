package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/12909 - 올바른 괄호

type stack []rune

func (s *stack) push(n rune) {
	*s = append(*s, n)
}
func (s *stack) pop() rune {
	if len(*s) == 0 {
		return -1
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func solution(str string) bool {
	s := make(stack, 0)

	for _, ch := range str {
		if ch == '(' {
			s.push(ch)
		} else {
			if s.pop() == -1 {
				return false
			}
		}
	}

	return len(s) == 0
}

func main() {
	case1 := solution("()()")
	fmt.Println(case1) // true
	case2 := solution("(())()")
	fmt.Println(case2) // true
	case3 := solution(")()(")
	fmt.Println(case3) // false
	case4 := solution("(()(")
	fmt.Println(case4) // false
}
