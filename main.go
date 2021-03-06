package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		printHelpAndExit()
	}

	fmt.Println(anagram(os.Args[1], os.Args[2]))
}

func printHelpAndExit() {
	fmt.Println("Wrong directive")
	fmt.Println("Usage: anagram word1 word2")
	os.Exit(1)
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var wg sync.WaitGroup
	m1, m2 := make(map[rune]int), make(map[rune]int)
	wg.Add(2)

	go func() {
		for _, r := range s1 {
			m1[r]++
		}

		wg.Done()
	}()

	go func() {
		for _, r := range s2 {
			m2[r]++
		}

		wg.Done()
	}()

	wg.Wait()

	for r, q := range m1 {
		if q != m2[r] {
			return false
		}
	}

	return true
}
