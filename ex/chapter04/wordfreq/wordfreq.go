package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func sortedKeys(M map[string]int) []string {
	keys := []string{}
	for key, _ := range M {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func wordFreq(r io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	M := make(map[string]int)
	for scanner.Scan() {
		M[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner.Scan: %v", err)
	}
	return M, nil
}

func main() {
	M, err := wordFreq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, key := range sortedKeys(M) {
		fmt.Printf("%s: %d\n", key, M[key])
	}
}
