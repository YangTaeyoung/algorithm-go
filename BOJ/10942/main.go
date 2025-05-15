package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/10942 팰린드롬?

func isPalindrome(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var (
		writer = bufio.NewWriter(os.Stdout)
		reader = bufio.NewReader(os.Stdin)
	)
	defer writer.Flush()

	// 1. 입력받기
	var N int
	if _, err := fmt.Fscan(reader, &N); err != nil {
		log.Fatal(err)
	}

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			log.Fatal(err)
		}
	}

	var M int
	if _, err := fmt.Fscan(reader, &M); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < M; i++ {
		var start, end int
		if _, err := fmt.Fscan(reader, &start, &end); err != nil {
			log.Fatal(err)
		}
		if isPalindrome(arr[start-1 : end]) {
			if _, err := fmt.Fprintln(writer, 1); err != nil {
				log.Fatal(err)
			}
		} else {
			if _, err := fmt.Fprintln(writer, 0); err != nil {
				log.Fatal(err)
			}
		}
	}
}
