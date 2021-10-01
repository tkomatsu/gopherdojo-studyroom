package main

import (
	"os"

	"jpg2png/convert"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		convert.Convert(path)
	}
}
