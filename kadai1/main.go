package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tkomatsu/gopherdojo-studyroom/kadai1/imageconv"
)

var src = flag.String("i", "jpg", "source image file format")
var dest = flag.String("o", "png", "destination image file format")

func main() {
	flag.Parse()
	paths := flag.Args()
	for _, path := range paths {
		doRecursive(path)
	}
}

func doRecursive(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", dir, err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			doRecursive(filepath.Join(dir, file.Name()))
		} else {
			err := imageconv.Convert(filepath.Join(dir, file.Name()), *dest, *src)
			if err != nil { fmt.Fprintln(os.Stderr, err) }
		}
	}
}
