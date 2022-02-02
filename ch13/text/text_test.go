package text

import (
	"testing"
)

func TestCoundCharacters(t *testing.T) {
	total, err := CountCharacters("testdata/sample1.txt")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if total != 35 {
		t.Error("Expcectd 35, got", total)
	}
	_, err = CountCharacters("testdata/no_file,txt")
	if err == nil {
		t.Error("Expected an error")
	}
}
