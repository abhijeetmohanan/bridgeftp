package utils

import "testing"

func TestParseInput(t *testing.T) {
	out, err := ParseInput("source", "sftp://user:password@host/myfile/out.txt")
	if err != nil {
		t.Error("failed parsing URL")
	}

	got := len(out)
	want := 5

	if got != want {
		t.Errorf("Wanted %d parameters git %d parameters ", want, got)
	}
}
