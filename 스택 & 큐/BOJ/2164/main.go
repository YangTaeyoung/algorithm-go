package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/2164 - 카드 2
type Queue struct {
	top   int
	queue []int
}

func (q *Queue) Len() int {
	return len(q.queue) - q.top
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(v int) {
	q.queue = append(q.queue, v)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	q.top++
	return q.queue[q.top-1]
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()
	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
		log.Fatal(err)
	}

	var q Queue
	for i := 1; i <= N; i++ {
		q.Push(i)
	}

	result := 0
	for !q.IsEmpty() {
		result = q.Pop()
		if found := q.Pop(); found != -1 {
			q.Push(found)
		}
	}
	if _, err := fmt.Fprintln(writer, result); err != nil {
		log.Fatal(err)
	}
}
