package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/10815 - 숫자 카드

func bs(arr []int, target int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			end = mid - 1
		}
		if arr[mid] < target {
			start = mid + 1
		}
	}

	return -1
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscanln(reader, &N); err != nil {
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
	searches := make([]int, M)
	for i := 0; i < M; i++ {
		if _, err := fmt.Fscan(reader, &searches[i]); err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(arr)
	for _, search := range searches {
		if found := bs(arr, search); found == -1 {
			if _, err := fmt.Fprint(writer, 0, " "); err != nil {
				log.Fatal(err)
			}
		} else {
			if _, err := fmt.Fprint(writer, 1, " "); err != nil {
				log.Fatal(err)
			}
		}
	}
}
