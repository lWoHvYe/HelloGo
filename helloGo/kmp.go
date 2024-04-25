package main

import "fmt"

func buildNext(pattern string) []int {
	m := len(pattern)
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func kmp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	pi := buildNext(pattern)
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	text := "ABCABCDABABCDABCDABDE"
	pattern := "ABCDABD"
	index := kmp(text, pattern)
	if index != -1 {
		fmt.Printf("Pattern found at index %d\n", index)
	} else {
		fmt.Println("Pattern not found")
	}
}
