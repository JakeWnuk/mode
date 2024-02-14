// Package utils contains functions for the main program
package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/net/html"
)

// ProcessFile reads the contents of a file and sends each line to the channel
//
// Args:
//
//	filename (string): The name of the file to read
//	ch (chan<- string): The channel to send the lines to
//	wg (*sync.WaitGroup): The WaitGroup to signal when done
//
// Returns:
//
//	None
func ProcessFile(filename string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(0)
	}

	buf, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(0)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				os.Exit(0)
			}
		}
	}()

	filescanner := bufio.NewScanner(buf)
	for filescanner.Scan() {
		ch <- filescanner.Text()
	}
}

// ProcessURL reads the contents of a URL and sends each sentence to the channel
//
// Args:
//
//	url (string): The URL to read
//	ch (chan<- string): The channel to send the sentences to
//	wg (*sync.WaitGroup): The WaitGroup to signal when done
//
// Returns:
//
//	None
func ProcessURL(url string, ch chan<- string, wg *sync.WaitGroup) {
	const maxRetries = 4
	defer wg.Done()

	var err error
	var resp *http.Response
	for attempts := 0; attempts <= maxRetries; attempts++ {

		resp, err = http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			time.Sleep(time.Second)
			continue
		}

		break
	}

	// Read Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	text := string(body)
	text = html.UnescapeString(text)
	var lines []string

	// Check the Content-Type of the response
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "text/html") {
		// Parse the HTML
		doc, err := html.Parse(strings.NewReader(text))
		if err != nil {
			panic(err)
		}

		// Traverse the HTML tree and extract the text
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.TextNode {
				lines = append(lines, n.Data)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	} else {
		sentences := strings.Split(text, "\n")
		for _, line := range sentences {
			lines = append(lines, line)
		}
	}

	// Iterate over the lines and split them
	for _, line := range lines {
		textMatch, _ := regexp.MatchString(`[^a-zA-Z0-9.,;:!?'"\- ]`, line)
		if strings.Contains(contentType, "text/html") {
			if textMatch {
				continue
			}
		} else {
			if !textMatch {
				continue
			}
		}

		sentences := strings.Split(line, ".")
		for _, sentence := range sentences {
			sentence = strings.TrimSpace(sentence)

			phrases := strings.Split(sentence, ",")
			for _, phrase := range phrases {
				if phrase != "" {
					ch <- phrase
				}
			}

			if sentence != "" {
				ch <- sentence
			}
		}
	}
}

// FilterTextForPrint processes the text and filters it based on the given parameters.
//
// Args:
//
//	text (string): The text to process
//	excludeTextLength (int): The minimum length of text to include
//	retainMap (map[string]bool): A map of words to retain
//	removeMap (map[string]bool): A map of words to remove
//
// Returns:
//
//	(string, error): The processed and filtered text, or an error if the text is not valid
func FilterTextForPrint(text string, excludeTextLength int, retainMap map[string]bool, removeMap map[string]bool) (string, error) {
	if len(text) <= excludeTextLength {
		return "", fmt.Errorf("text length is less than or equal to: %d", excludeTextLength)
	}

	if len(retainMap) == 0 && len(removeMap) == 0 {
		return text, nil
	}

	if _, ok := retainMap[text]; ok || len(retainMap) == 0 {
		if _, ok := removeMap[text]; !ok || len(removeMap) == 0 {
			return text, nil
		}
	}

	return "", fmt.Errorf("text is not in the retain list")
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

// PrintItems prints an array of items to stdout
//
// Args:
//
//	items ([]string): The array of items to print
//	freq (map[string]int): A map of item frequencies
//	verbose (bool): If true, the count is printed
//	min (int): Minimum frequency of an item to print
//
// Returns:
//
//	None
func PrintItems(items []string, freq map[string]int, verbose bool, min int) {
	count := 0
	for _, item := range items {
		if freq[item] >= min {
			if verbose {
				fmt.Printf("%d %s\n", freq[item], item)
			} else {
				fmt.Println(item)
			}
			count++
		}
	}

}

// LoadArgumentFileContents reads the contents of the given files and returns a map of words
//
// Args:
//
//	filenames ([]string): The names of the files to read
//
// Returns:
//
//	(map[string]bool): A map of words from the files
func LoadArgumentFileContents(filenames []string) map[string]bool {
	wordMap := make(map[string]bool)
	for _, filename := range filenames {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		fileWords := strings.Split(string(data), "\n")
		for _, word := range fileWords {
			wordMap[word] = true
		}
	}
	return wordMap
}

// TokenizeParse parses the tokens and returns a slice of words
//
// Args:
//
//	str (string): The string to parse
//
//	Returns:
//	[]string: A slice of words
//	error: Any error that occurred
func TokenizeParse(str string) ([]string, error) {
	// Store word as is
	returnSlice := make([]string, 0)
	returnSlice = append(returnSlice, str)

	// Capitalize after space
	str = capitalizeAfterSpace(str)
	returnSlice = append(returnSlice, str)

	// Pop spaces and all punctuation
	punctReplacer := strings.NewReplacer(" ", "", ".", "", ",", "", ";", "", ":", "", "!", "", "?", "", "'", "", "\"", "", "-", "")
	str = punctReplacer.Replace(str)
	returnSlice = append(returnSlice, str)

	// Lowercase
	str = strings.ToLower(str)
	returnSlice = append(returnSlice, str)

	return returnSlice, nil
}

func capitalizeAfterSpace(str string) string {
	// Convert string to a rune slice for efficient character manipulation
	chars := []rune(str)

	// Iterate through the characters
	for i, char := range chars {
		// Capitalize the first character or any character after a space
		if i == 0 || unicode.IsSpace(chars[i-1]) {
			chars[i] = unicode.ToUpper(char)
		}
	}

	// Return the capitalized string
	return string(chars)
}
