package main

import (
	"bufio"
	"fmt"
	"os"
)

// Originally from https://stackoverflow.com/a/30230552
func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig []rune, p []int) []rune {
	result := append([]rune{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

// parseWords parses the given path to a map of words.
func parseWords(path string) map[string]bool {
	result := make(map[string]bool, 0)

	r, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		os.Exit(1)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		word := scanner.Text()
		result[word] = true
	}

	return result
}

func main() {
	if len(os.Args) < 2 { //nolint:gomnd
		fmt.Println("Error: need a word.  Usage: descrambluh <someword>")
	}
	word := os.Args[1]
	wordPath := "/usr/share/dict/words"
	words := parseWords(wordPath)

	orig := []rune(word)
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		permutation := getPerm(orig, p)

		s := string(permutation)
		if _, ok := words[s]; ok {
			fmt.Println(s, "was found")
		}
	}
}