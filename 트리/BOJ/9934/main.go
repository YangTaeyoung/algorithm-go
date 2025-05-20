package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// https://www.acmicpc.net/problem/9934 - 완전 이진 트리

// 더미 트리를 만들고 순회 후 TraversalOrder를 만들어서 기존 배열 출력 시 해당 Order에 맞게 출력함으로써 해결함

func InOrder(orders *[]int, trees []int, current int) {
	left := current*2 + 1
	right := current*2 + 2

	if len(trees)-1 >= left {
		InOrder(orders, trees, left)
	}
	*orders = append(*orders, current)
	if len(trees)-1 >= right {
		InOrder(orders, trees, right)
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var K int
	if _, err := fmt.Fscanln(reader, &K); err != nil {
		log.Fatal(err)
	}

	length := int(math.Pow(2, float64(K))) - 1
	tree := make([]int, length)
	orders := make([]int, length)

	for i := 0; i < length; i++ {
		if _, err := fmt.Fscan(reader, &orders[i]); err != nil {
			log.Fatal(err)
		}
	}

	traversalOrders := make([]int, 0, length)
	InOrder(&traversalOrders, tree, 0)

	for i, traversalOrder := range traversalOrders {
		tree[traversalOrder] = orders[i]
	}

	for i := 0; i < K; i++ {
		start := int(math.Pow(2, float64(i))) - 1
		end := int(math.Pow(2, float64(i+1))) - 1
		for j := start; j < end; j++ {
			fmt.Print(tree[j], " ")
		}
		fmt.Println()
	}
}
