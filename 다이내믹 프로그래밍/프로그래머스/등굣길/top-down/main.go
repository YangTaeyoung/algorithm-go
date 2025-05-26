package main

import "fmt"

func dp(m int, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d %d", m, n)
	if value, exists := memo[key]; exists {
		return value
	}

	if m == 1 {
		memo[key] = dp(m, n-1, memo)
	} else if n == 1 {
		memo[key] = dp(m-1, n, memo)
	} else {
		memo[key] = dp(m, n-1, memo) + dp(m-1, n, memo)
	}

	return memo[key] % 1000000007
}

func solution(m int, n int, puddles [][]int) int {
	var memo = make(map[string]int)

	memo["1 1"] = 1
	memo["1 2"] = 1
	memo["2 1"] = 1
	for _, puddle := range puddles {
		key := fmt.Sprintf("%d %d", puddle[0], puddle[1])
		memo[key] = 0
	}

	answer := dp(m, n, memo)

	return answer
}

func main() {
	type testCase struct {
		name    string
		m, n    int
		puddles [][]int
		want    int
	}

	tests := []testCase{
		{"1x1, no puddle", 1, 1, [][]int{}, 1},
		{"2x2, no puddle", 2, 2, [][]int{}, 2},
		{"3x3, no puddle", 3, 3, [][]int{}, 6},
		{"4x3, one puddle at (2,2)", 4, 3, [][]int{{2, 2}}, 4},
		{"3x3, puddle blocking first row", 3, 3, [][]int{{2, 1}}, 3},
		{"3x3, puddle blocking last column path", 3, 3, [][]int{{3, 2}}, 3},
		{"5x5, no puddle", 5, 5, [][]int{}, 70},
		{"100x1, long corridor", 100, 1, [][]int{}, 1},

		// 엣지/극단 케이스
		{"blocked around start", 3, 3, [][]int{{2, 1}, {1, 2}}, 0},
		{"blocked around end", 3, 3, [][]int{{2, 3}, {3, 2}}, 0},
		{"2x2, full block", 2, 2, [][]int{{2, 1}, {1, 2}}, 0},
		{"1x100, puddle in middle", 1, 100, [][]int{{1, 50}}, 0},
		{"100x1, puddle in middle", 100, 1, [][]int{{50, 1}}, 0},
		{"100x100, no puddle", 100, 100, [][]int{}, 690285631},
		{"10x10, no puddle", 10, 10, [][]int{}, 48620},
	}

	pass, fail := 0, 0
	for _, tc := range tests {
		got := solution(tc.m, tc.n, tc.puddles)
		if got == tc.want {
			fmt.Printf("[PASS] %s: got=%d\n", tc.name, got)
			pass++
		} else {
			fmt.Printf("[FAIL] %s: got=%d, want=%d\n", tc.name, got, tc.want)
			fail++
		}
	}
	fmt.Printf("\nTotal: %d passed, %d failed\n", pass, fail)
}
