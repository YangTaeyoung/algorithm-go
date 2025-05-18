package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/2178 - 미로찾기

type Coordinate struct {
	X, Y int
}

type Coordinates []Coordinate

func (c *Coordinates) isEmpty() bool {
	return len(*c) == 0
}
func (c *Coordinates) Pop() Coordinate {
	result := (*c)[0]
	*c = (*c)[1:]

	return result
}

func adjacent(graph [][]int, c Coordinate) Coordinates {
	result := make(Coordinates, 0)
	if c.X+1 < len(graph) {
		if graph[c.X+1][c.Y] == 1 {
			result = append(result, Coordinate{
				X: c.X + 1,
				Y: c.Y,
			})
		}
	}
	if c.Y+1 < len(graph[c.X]) {
		if graph[c.X][c.Y+1] == 1 {
			result = append(result, Coordinate{
				X: c.X,
				Y: c.Y + 1,
			})
		}
	}
	if c.X-1 >= 0 {
		if graph[c.X-1][c.Y] == 1 {
			result = append(result, Coordinate{
				X: c.X - 1,
				Y: c.Y,
			})
		}
	}
	if c.Y-1 >= 0 {
		if graph[c.X][c.Y-1] == 1 {
			result = append(result, Coordinate{
				X: c.X,
				Y: c.Y - 1,
			})
		}
	}

	return result
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M int
	if _, err := fmt.Fscan(reader, &N, &M); err != nil {
		log.Fatalln(err)
	}

	graph := make([][]int, N)
	visited := make([][]int, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]int, M)
		graph[i] = make([]int, M)
		var row string
		if _, err := fmt.Fscan(reader, &row); err != nil {
			log.Fatalln(err)
		}
		for j := range row {
			graph[i][j] = int(row[j] - '0')
		}
	}

	q := Coordinates{
		{X: 0, Y: 0},
	}
	visited[0][0] = 1

	var current Coordinate
	for !q.isEmpty() {
		current = q.Pop()
		if current.X == N-1 && current.Y == M-1 {
			break
		}
		for _, coordinate := range adjacent(graph, current) {
			if visited[coordinate.X][coordinate.Y] == 0 {
				q = append(q, coordinate)
				visited[coordinate.X][coordinate.Y] = visited[current.X][current.Y] + 1
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", visited[N-1][M-1])
}
