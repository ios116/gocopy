package main

import (
	"flag"
	"fmt"
	"github.com/ios116/gocopy"
	"log"
	"os"
)

var from string
var to string
var limit int64
var offset int64
var bs int64

func init() {
	flag.StringVar(&from, "from", "", "input file")
	flag.StringVar(&to, "to", "", "output file")
	flag.Int64Var(&limit, "limit", 1024, "bytes limit")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
	flag.Int64Var(&bs, "bs", 1, "Set output block size to n bytes")
}

func main() {
	flag.Parse()

	sours, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}

	gocopy := &gocopy.GoCopy{
		R:      sours,
		W:      out,
		Limit:  limit,
		Bs:     bs,
		Offset: offset,
	}
	n, err := gocopy.Copier()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied %d bytes\n\n", n)
}
