package utils

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestAddNgramCount(t *testing.T) {
	freq := make(map[string]int)
	item := "the quick brown fox jumps over the lazy dog"
	expected := map[string]int{
		"the": 2, "quick": 1, "brown": 1, "fox": 1,
		"jumps": 1, "over": 1, "lazy": 1, "dog": 1,
	}
	result := AddNgramCount(item, freq)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAddSegmentCount(t *testing.T) {
	freq := make(map[string]int)
	item := "thequickbrownfoxjumpsoverthelazydog"
	dict := []string{"quick", "lazy", "fox"}
	expected := map[string]int{
		"quick": 1,
	}
	result := AddSegmentCount(item, freq, dict)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestCountFrequencies(t *testing.T) {
	input := "the quick brown fox jumps over the lazy dog\nthe quick brown fox jumps over the lazy dog\n"
	stdIn := bufio.NewScanner(strings.NewReader(input))
	split := true
	file := ""
	freq := make(map[string]int)
	expected := map[string]int{
		"the": 4, "quick": 2, "brown": 2, "fox": 2,
		"jumps": 2, "over": 2, "lazy": 2, "dog": 2,
		"the quick brown fox jumps over the lazy dog": 2,
	}
	result := CountFrequencies(stdIn, split, file, freq)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortItems(t *testing.T) {
	freq := map[string]int{
		"the": 9, "quick": 10, "brown": 7, "fox": 6,
		"jumps": 5, "over": 4, "lazy": 3, "dog": 2,
		"the quick brown fox jumps over the lazy dog": 1,
	}
	expected := []string{
		"quick", "the", "brown", "fox",
		"jumps", "over", "lazy", "dog",
		"the quick brown fox jumps over the lazy dog",
	}
	result := SortItems(freq)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSegmentWords(t *testing.T) {
	input := "thequickbrownfoxjumpsoverthelazydog"
	dict := []string{"quick", "lazy", "fox"}
	expected := []string{"quick"}
	result := SegmentWords(input, dict)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
