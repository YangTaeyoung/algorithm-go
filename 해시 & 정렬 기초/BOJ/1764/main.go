package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://www.acmicpc.net/problem/1764 - 듣보잡

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M int
	if _, err := fmt.Fscan(reader, &N, &M); nil != err {
		log.Fatal(err)
	}

	names := make(map[string]bool)
	for i := 0; i < N; i++ {
		var name string
		if _, err := fmt.Fscan(reader, &name); nil != err {
			log.Fatal(err)
		}
		names[name] = true
	}

	var answer []string
	for i := 0; i < M; i++ {
		var name string
		if _, err := fmt.Fscan(reader, &name); nil != err {
			log.Fatal(err)
		}
		if names[name] {
			answer = append(answer, name)
		}
	}

	sort.Strings(answer)

	if _, err := fmt.Fprintln(writer, len(answer)); err != nil {
		log.Fatal(err)
	}

	for _, name := range answer {
		if _, err := fmt.Fprintln(writer, name); err != nil {
			log.Fatal(err)
		}
	}

}
