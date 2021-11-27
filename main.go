package main

import (
	"flag"
	"os"
)

func main() {

	n := flag.Int("n", 10, "Number of records")
	flag.Parse()

	g := NewGenerator(*n, "GB")
	g.Create(os.Stdout)
}
