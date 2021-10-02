package main

import (
	"flag"

	"github.com/tkomatsu/gopherdojo-studyroom/kadai1/imageconv"
)

var srcType = flag.String("f", "jpeg", "source image file format")

func main() {
	flag.Parse()
	paths := flag.Args()
	for _, path := range paths {
		imageconv.Convert(path, *srcType)
	}
}
