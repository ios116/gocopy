package copier

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGoCopy(t *testing.T) {

	sours, err := os.Open("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	dest, err := os.Create("./out.txt")
	if err != nil {
		t.Fatal(err)
	}

	gocopy := &GoCopy{
		R:      sours,
		W:      dest,
		Limit:  5,
		Bs:     1,
		Offset: 3,
	}
	_, err = gocopy.Copier()
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile("./out.txt")
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "gramm" {
		t.Fatalf("we have %s must be %s", string(b), "gramm")
	}
}
