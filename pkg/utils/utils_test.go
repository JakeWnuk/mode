package utils

import (
	"reflect"
	"testing"
)

func TestFilterTextForPrint(t *testing.T) {
	tests := []struct {
		text              string
		excludeTextLength int
		retainMap         map[string]bool
		removeMap         map[string]bool
		want              string
		wantErr           bool
	}{
		{"hello", 3, map[string]bool{}, map[string]bool{}, "hello", false},
		{"hi", 3, map[string]bool{}, map[string]bool{}, "", true},
		{"hello", 3, map[string]bool{"hello": true}, map[string]bool{}, "hello", false},
		{"hi", 3, map[string]bool{"hello": true}, map[string]bool{}, "", true},
		{"hello", 3, map[string]bool{}, map[string]bool{"hello": true}, "", true},
	}

	for _, tt := range tests {
		got, err := FilterTextForPrint(tt.text, tt.excludeTextLength, tt.retainMap, tt.removeMap)
		if (err != nil) != tt.wantErr {
			t.Errorf("FilterTextForPrint(%q, %d, %v, %v) error = %v, wantErr %v", tt.text, tt.excludeTextLength, tt.retainMap, tt.removeMap, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("FilterTextForPrint(%q, %d, %v, %v) = %v, want %v", tt.text, tt.excludeTextLength, tt.retainMap, tt.removeMap, got, tt.want)
		}
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
