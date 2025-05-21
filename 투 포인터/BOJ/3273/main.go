package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sort.Ints(arr)

	count := 0
	start, end := 0, len(arr)-1
	for start < end {
		sum := arr[start] + arr[end]
		if sum == target {
			count++
			start++
			end--
		}
		if sum > target {
			end--
		}
		if sum < target {
			start++
		}
	}

	fmt.Println(count)
}
