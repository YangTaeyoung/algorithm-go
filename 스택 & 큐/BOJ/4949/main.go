package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/4949 - 균형잡힌 세상
type Stack []rune

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var (
		sentences []string
		err       error
	)
	for {
		var sentence string
		if sentence, err = reader.ReadString('\n'); err != nil {
			log.Fatal(err)
		}
		if sentence == ".\n" {
			break
		}
		sentences = append(sentences, sentence)
	}

	for _, sentence := range sentences {
		stack := make(Stack, 0)
		result := "yes"
		for _, r := range sentence {
			if r == '(' || r == '[' {
				stack.Push(r)
			}
			if r == ')' {
				if stack.Pop() != '(' {
					result = "no"
					break
				}
			}
			if r == ']' {
				if stack.Pop() != '[' {
					result = "no"
					break
				}
			}
		}
		if len(stack) > 0 {
			result = "no"
		}

		if _, err = fmt.Fprintln(writer, result); err != nil {
			log.Fatal(err)
		}
	}

}
