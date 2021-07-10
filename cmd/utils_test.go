package cmd

import (
	"fmt"
	"testing"
)

func Test_mostFrequentExt(t *testing.T) {
	mapExt := map[string]int{".md": 5, ".txt": 5, ".cer": 10, ".ebv": 2}

	result := mostFrequentExt(mapExt)
	fmt.Println("map to test  Frequent Ext :", mapExt)
	fmt.Println("Test_mostFrequentExt result :", result)
	if result.Extension != ".cer" && result.NumOccurrences != 10 {
		t.Fatal("expected {.cer 10} got ", result)
	}
}
