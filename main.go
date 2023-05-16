// Package that contains the primary logic for mode and the CLI
package main

import (
	"flag"
	"fmt"

	"github.com/jakewnuk/mode/pkg/utils"
)

var version = "1.0.0"

func main() {
	count := flag.Bool("c", false, "Display the frequency count of each item\nExample: mode -c")
	ngrams := flag.Bool("s", false, "Split items into n-grams by spaces and add to count.\nExample: mode -s")
	dict := flag.String("f", "", "Parse items from a dictionary file and add to count. The file should contain one item per line.\nExample: mode -f dict.txt")
	exclude := flag.Int("x", 0, "Exclude items below a length from output. Length should be an integer.\nExample: mode -x 5")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of mode (version %s):\n", version)
		flag.PrintDefaults()
	}

	flag.Parse()

	freq := make(map[string]int)
	scanner, err := utils.ReadStdin()
	utils.CheckError(err)

	freq = utils.CountFrequencies(scanner, *ngrams, *dict, freq)
	items := utils.SortItems(freq)
	utils.PrintItems(items, freq, *count, *exclude)
}
