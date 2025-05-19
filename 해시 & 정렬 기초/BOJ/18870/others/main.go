package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 내 풀이와 방식 자체는 동일하지만 Scan과 동시에 map을 만들고, map을 그대로 활용해서 RankMap까지 만드는게 인상 깊음
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	x := make([]int, n)
	numSet := make(map[int]int)
	for idx := 0; idx < n; idx++ {
		scanner.Scan()
		x[idx], _ = strconv.Atoi(scanner.Text())
		numSet[x[idx]] = -1
	}

	sortedKeys := make([]int, len(numSet))
	idx := 0
	for k := range numSet {
		sortedKeys[idx] = k
		idx++
	}
	sort.Ints(sortedKeys)
	for idx, key := range sortedKeys {
		numSet[key] = idx
	}

	for idx, val := range x {
		x[idx] = numSet[val]
	}

	result := make([]string, n)
	for idx, val := range x {
		result[idx] = strconv.Itoa(val)
	}

	fmt.Println(strings.Join(result, " "))
}
