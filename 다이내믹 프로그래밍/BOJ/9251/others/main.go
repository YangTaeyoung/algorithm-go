// Bottom - Up
package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

func main() {
	defer bw.Flush()
	A, B := input(), input()
	DP := make([]int, len(B))
	for i := 0; i < len(A); i++ {
		cnt := 0
		for j := 0; j < len(B); j++ {
			if cnt < DP[j] {
				cnt = DP[j]
			} else if A[i] == B[j] {
				DP[j] = cnt + 1
			}
		}
	}
	bw.WriteString(strconv.Itoa(slices.Max(DP)))
}

var (
	bs = bufio.NewScanner(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func input() string {
	bs.Scan()
	return bs.Text()
}
