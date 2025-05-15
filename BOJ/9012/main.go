package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/9012 - 괄호
func main() {
	var (
		writer = bufio.NewWriter(os.Stdout)
		reader = bufio.NewReader(os.Stdin)
	)
	defer writer.Flush()
	var T int
	if _, err := fmt.Fscan(reader, &T); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < T; i++ {
		var (
			str    string
			stack  string
			result string
		)
		if _, err := fmt.Fscan(reader, &str); err != nil {
			log.Fatal(err)
		}

		for _, s := range str {
			if s == '(' {
				stack += string(s)
				continue
			}

			if len(stack) == 0 {
				result = "NO"
				break
			}

			stack = stack[:len(stack)-1]
		}

		if result == "" {
			if len(stack) > 0 {
				result = "NO"
			} else {
				result = "YES"
			}
		}

		if _, err := fmt.Fprintln(writer, result); err != nil {
			log.Fatal(err)
		}
	}
}
