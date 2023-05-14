// Package utils contains functions for the main mode program
package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// AddNgramCount splits input by spaces into ngrams then adds their frequency
func AddNgramCount(item string, freq map[string]int) map[string]int {
	for _, ngram := range strings.Fields(item) {
		freq[ngram]++
	}
	return freq
}

// AddSegmentCount parses input by a dict into segments then adds their frequency
func AddSegmentCount(item string, freq map[string]int, dict []string) map[string]int {
	segments := SegmentWords(item, dict)
	for _, tk := range segments {
		freq[tk]++
	}
	return freq
}

// CountFrequencies takes stdin and returns a map of item frequencies
// If split is provided input is split into ngrams by strings.Fields()
// If parse is provided input is parsed into segments by a dict
// Items are added to the base frequency which is stdin input
func CountFrequencies(stdIn *bufio.Scanner, split bool, file string, freq map[string]int) map[string]int {
	var dict []string
	parse := false

	if file != "" {
		parse = true
		dict = LoadDict(file)
	}

	for stdIn.Scan() {
		item := stdIn.Text()

		if split && parse {
			freq = AddNgramCount(item, freq)
			freq = AddSegmentCount(item, freq, dict)
		} else if split {
			freq = AddNgramCount(item, freq)
		} else if parse {
			freq = AddSegmentCount(item, freq, dict)
		}

		freq[item]++
	}

	return freq
}

// SortItems sorts items in the map by frequency and returns a sorted slice
func SortItems(freq map[string]int) []string {
	items := make([]string, 0, len(freq))

	for item := range freq {
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return freq[items[i]] > freq[items[j]]
	})

	return items
}

// PrintItems prints an array to stdout
// If verbose is provided the count is printed
func PrintItems(items []string, freq map[string]int, verbose bool) {
	for _, item := range items {
		if verbose {
			fmt.Printf("%d %s\n", freq[item], item)
		} else {
			fmt.Println(item)
		}
	}
}

// SegmentWords splits an input string into substrings that match a dict
// NOTE: Attempts to match largest substrings first then stops search after match
func SegmentWords(input string, words []string) []string {
	var result []string

	for i := 0; i < len(input); {
		matched := false
		for _, word := range words {
			if len(word) <= len(input)-i && input[i:i+len(word)] == word {
				result = append(result, word)
				i += len(word)
				matched = true
				break
			}
		}
		// No match then increment until a match is found
		if !matched {
			i++
		} else {
			break
		}
	}

	return result
}

// ReadStdin reads stdin to a scanner object and returns the scanner and error
func ReadStdin() (*bufio.Scanner, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return scanner, nil
}

// LoadDict loads a dictionary of words from a file
func LoadDict(filename string) []string {
	file, err := os.Open(filename)
	CheckError(err)
	defer file.Close()

	var dict []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		CheckError(err)
	}

	// Sort by length in descending order
	// NOTE: Needed for SegmentWords and is currently the only calling function
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) > len(dict[j])
	})

	return dict
}

// CheckError is a general error handler
func CheckError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(0)
	}
}
