package main

import "fmt"

func min(arr ...int) int {
	value := arr[0]
	for _, n := range arr {
		if n < value {
			value = n
		}
	}
	return value
}

func solution(name string) int {
	verticalMove := 0
	for _, alpha := range name {
		asc := int(alpha - 'A')
		desc := int('Z' - alpha + 1)
		verticalMove += min(asc, desc)
	}

	horizontalMove := len(name) - 1
	for i := 0; i < len(name)-1; i++ {
		next := i + 1
		for next < len(name) && name[next] == 'A' {
			next++
		}
		// 오른쪽으로 갔다가 왼쪽으로 돌아와 왼쪽으로 이동하는 경우
		moveRightLeft := (i * 2) + (len(name) - next)
		// 왼쪽으로 갔다가 오른쪽으로 돌아와 오른쪽으로 이동하는 경우
		moveLeftRight := (len(name)-next)*2 + i

		horizontalMove = min(horizontalMove, moveRightLeft, moveLeftRight)
	}

	return verticalMove + horizontalMove
}

func main() {
	case1 := solution("JEROEN")
	fmt.Println(case1)
	case2 := solution("JAN")
	fmt.Println(case2)
}
