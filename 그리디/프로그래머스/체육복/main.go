package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, value int) int {
	minIdx := 0
	maxIdx := len(arr) - 1

	for minIdx <= maxIdx {
		mid := (minIdx + maxIdx) / 2
		if arr[mid] == value {
			return mid
		}
		if arr[mid] < value {
			minIdx = mid + 1
		}
		if arr[mid] > value {
			maxIdx = mid - 1
		}
	}
	return -1
}

func removeDuplicates(arr1 []int, arr2 []int) ([]int, []int) {
	result1 := make([]int, 0, len(arr1))

	for i, num := range arr1 {
		if found := binarySearch(arr2, num); found != -1 {
			continue
		}
		result1 = append(result1, arr1[i])
	}

	result2 := make([]int, 0, len(arr2))
	for i, num := range arr2 {
		if found := binarySearch(arr1, num); found != -1 {
			continue
		}
		result2 = append(result2, arr2[i])
	}

	return result1, result2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/42862 - 체육복
func solution(n int, lost []int, reserve []int) int {
	sort.Ints(lost)
	sort.Ints(reserve)
	lost, reserve = removeDuplicates(lost, reserve)

	answer := n - len(lost)
	for _, r := range reserve {
		if found := binarySearch(lost, r-1); found != -1 { // 자신보다 이전 번호에서 여분이 있음.
			lost = append(lost[:found], lost[found+1:]...) // 여분 삭제
			answer++
			continue
		}
		if found := binarySearch(lost, r+1); found != -1 { // 자신보다 이후 번호에서 여분이 있음.
			lost = append(lost[:found], lost[found+1:]...) // 여분 삭제
			answer++
			continue
		}
	}
	return answer
}

func main() {
	case1 := solution(5, []int{2, 4}, []int{1, 3, 5})
	fmt.Println(case1)
	case2 := solution(5, []int{2, 4}, []int{3})
	fmt.Println(case2)
	case3 := solution(3, []int{3}, []int{1})
	fmt.Println(case3)
	case4 := solution(4, []int{2, 3, 4}, []int{1, 2, 3})
	fmt.Println(case4)
}
