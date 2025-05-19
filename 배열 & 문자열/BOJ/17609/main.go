package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.acmicpc.net/problem/17609 - 회문

func isPseudoPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isPalindrome(s string) int {
	start := 0
	end := len(s) - 1
	result := 0
	for start < end {
		if s[start] != s[end] {
			if isPseudoPalindrome(s, start+1, end) {
				return 1
			}
			if isPseudoPalindrome(s, start, end-1) {
				return 1
			}
			return 2
		}
		start++
		end--
	}
	return result
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var T int
	if _, err := fmt.Fscanln(reader, &T); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		var sentence string
		if _, err := fmt.Fscan(reader, &sentence); err != nil {
			log.Fatal(err)
		}

		if _, err := fmt.Fprintln(writer, isPalindrome(sentence)); err != nil {
			log.Fatal(err)
		}
	}
}
