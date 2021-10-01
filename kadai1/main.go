package main

import (
	"flag"

	"convert/convert"
)

var srcType = flag.String("f", "jpeg", "source image file format")

func main() {
	flag.Parse()
	paths := flag.Args()
	for _, path := range paths {
		convert.Convert(path)
	}
}
