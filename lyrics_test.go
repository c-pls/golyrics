package golyrics

import (
	"strings"
	"testing"
)

func TestLyrics(t *testing.T) {
	res := GetLyrics("In My Life", "The Beatles")
	want := "There is no one compares with you"

	if check := strings.Contains(res, want); !check {
		t.Error()
	}

}
