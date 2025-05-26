package main

import "slices"

func canHandled(times []int, time int64, human int) bool {
	var total int64 = 0
	for _, t := range times {
		total += time / int64(t)
	}
	return total >= int64(human)
}

func solution(n int, times []int) int64 {
	maxTime := int64(slices.Max(times)) * int64(n)
	minTime := int64(slices.Min(times))

	answer := int64(0)
	for minTime <= maxTime {
		mid := (minTime + maxTime) / 2
		if canHandled(times, mid, n) {
			maxTime = mid - 1
			answer = mid
		} else {
			minTime = mid + 1
		}
	}

	return answer
}

func main() {
	case1 := solution(6, []int{7, 10})
	println(case1) // 28
	case2 := solution(5, []int{2, 3, 4})
	println(case2) // 6
}
