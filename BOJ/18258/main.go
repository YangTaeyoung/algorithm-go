package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/18258 - 큐2
func main() {
	var (
		writer = bufio.NewWriter(os.Stdout)
		reader = bufio.NewReader(os.Stdin)
	)
	defer writer.Flush()

	var N int
	if _, err := fmt.Fscan(reader, &N); err != nil {
		log.Fatal(err)
	}

	var q []int

	for i := 0; i < N; i++ {
		var command string
		if _, err := fmt.Fscan(reader, &command); err != nil {
			log.Fatal(err)
		}

		switch command {
		case "push":
			var v int
			if _, err := fmt.Fscan(reader, &v); err != nil {
				log.Fatal(err)
			}
			q = append(q, v)
		case "front":
			var top int
			if len(q) == 0 {
				top = -1
			} else {
				top = q[0]
			}

			if _, err := fmt.Fprintln(writer, top); err != nil {
				log.Fatal(err)
			}
		case "back":
			var back int
			if len(q) == 0 {
				back = -1
			} else {
				back = q[len(q)-1]
			}

			if _, err := fmt.Fprintln(writer, back); err != nil {
				log.Fatal(err)
			}
		case "pop":
			var top int
			if len(q) == 0 {
				top = -1
			} else {
				top = q[0]
				q = q[1:]
			}

			if _, err := fmt.Fprintln(writer, top); err != nil {
				log.Fatal(err)
			}
		case "size":
			if _, err := fmt.Fprintln(writer, len(q)); err != nil {
				log.Fatal(err)
			}
		case "empty":
			if len(q) == 0 {
				if _, err := fmt.Fprintln(writer, 1); err != nil {
					log.Fatal(err)
				}
			} else {
				if _, err := fmt.Fprintln(writer, 0); err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}
