// Package main controls the main program
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/jakewnuk/mode/pkg/utils"
)

var version = "0.0.1"

type argumentFilesFlag []string

func (w *argumentFilesFlag) String() string {
	return fmt.Sprint(*w)
}

func (w *argumentFilesFlag) Set(value string) error {
	*w = append(*w, value)
	return nil
}

func main() {
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	var retain argumentFilesFlag
	var remove argumentFilesFlag
	ch := make(chan string)
	contentMap := make(map[string]int)
	urlRegex := `^((http|https)://)([\w-]+\.)+[\w-]+(/[\w- ;,./?%&=]*)?$`

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Mode version (%s):\n\n", version)
		fmt.Fprintf(os.Stderr, "mode [options] [URLS/FILES] [...]\nAccepts standard input and/or additonal arguments.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	count := flag.Bool("c", false, "Show the frequency count of each item")
	parse := flag.Bool("a", false, "Perform additional parsing of each item")
	minimum := flag.Int("m", 0, "Minimum frequency to include in output. Value should be an integer.")
	exclude := flag.Int("x", 0, "Exclude items less than or equal to a length from output. Length should be an integer.")
	flag.Var(&retain, "w", "Only include items in a file.")
	flag.Var(&remove, "v", "Only include items not in a file.")
	flag.Parse()

	// Parse any retain/remove files
	retainMap := utils.LoadArgumentFileContents(retain)
	removeMap := utils.LoadArgumentFileContents(remove)

	// Parse non-flag arguments into channels
	for _, arg := range flag.Args() {
		wg.Add(1)
		match, _ := regexp.MatchString(urlRegex, arg)
		if match {
			go utils.ProcessURL(arg, ch, &wg)
		} else {
			go utils.ProcessFile(arg, ch, &wg)
		}
	}

	// Check for standard input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		wg.Add(1)
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				ch <- scanner.Text()
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Process gathered text
	for line := range ch {
		mutex.Lock()
		// Split by spaces and add to map
		lines := strings.Fields(line)
		// Add the line itself to the map
		lines = []string{line}

		for _, line := range lines {
			str, err := utils.FilterTextForPrint(line, *exclude, retainMap, removeMap)
			if err == nil {
				var tokens []string
				if *parse {
					tokens, err = utils.TokenizeParse(str)
				} else {
					tokens = append(tokens, str)
				}
				if err != nil {
					continue
				}
				for _, token := range tokens {
					contentMap[token]++
				}
			}
		}
		mutex.Unlock()
	}
	items := utils.SortItems(contentMap)
	utils.PrintItems(items, contentMap, *count, *minimum)
}
