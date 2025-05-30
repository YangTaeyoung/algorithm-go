package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/3273 - 두 수의 합

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscanln(reader, &n); err != nil {
		log.Fatal(err)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := fmt.Fscanln(reader); err != nil {
		log.Fatal(err)
	}
	var target int
	if _, err := fmt.Fscanln(reader, &target); err != nil {
		log.Fatal(err)
	}

	count := 0
	idxMap := make(map[int]int)
	for i, num := range arr {
		if _, ok := idxMap[target-num]; ok {
			count++
		}
		idxMap[num] = i
	}

	if _, err := fmt.Fprintln(writer, count); err != nil {
		log.Fatal(err)
	}
}
