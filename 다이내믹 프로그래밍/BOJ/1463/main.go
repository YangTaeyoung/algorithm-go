package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/1463 - 1로 만들기
var memo = make(map[int]int)

func Compute(num int) int {
	if value, exists := memo[num]; exists {
		return value
	}
	if num == 1 {
		return 0
	}

	best := Compute(num-1) + 1
	if num%3 == 0 {
		if t := Compute(num/3) + 1; t < best {
			best = t
		}
	}
	if num%2 == 0 {
		if t := Compute(num/2) + 1; t < best {
			best = t
		}
	}
	memo[num] = best

	return best
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintln(writer, Compute(N))
}
