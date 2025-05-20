package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 다른 풀이를 보고 다른 방식으로 풀어봄
// ChatGPT의 설명을 받음: https://chatgpt.com/share/682c7096-86dc-800d-8957-702f8377435a
// InOrder의 인덱스 특성을 이용하여 해결함
func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var K int
	if _, err := fmt.Fscanln(reader, &K); err != nil {
		log.Fatal(err)
	}

	length := 1<<K - 1
	inOrder := make([]int, length)

	for i := 0; i < length; i++ {
		if _, err := fmt.Fscan(reader, &inOrder[i]); err != nil {
			log.Fatal(err)
		}
	}

	idxs := []int{length / 2}

	for depth := 1; depth < K; depth++ {
		children := make([]int, 0, 1<<(depth))
		interval := 1 << (K - depth - 1)
		for _, idx := range idxs {
			_, _ = fmt.Fprint(writer, inOrder[idx], " ")
			children = append(children, idx-interval, idx+interval)
		}
		idxs = children
		_, _ = fmt.Fprintln(writer)
	}

	for _, idx := range idxs {
		_, _ = fmt.Fprint(writer, inOrder[idx], " ")
	}
}
