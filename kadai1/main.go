package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"path/filepath"
	"os"
	"strings"

	"github.com/tkomatsu/gopherdojo-studyroom/kadai1/imageconv"
)

var srcType = flag.String("f", "jpeg", "source image file format")

func main() {
	flag.Parse()
	paths := flag.Args()
	for _, path := range paths {
		doRecursive(path, *srcType)
	}
}

func doRecursive(dir string, srcType string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", dir, err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			doRecursive(filepath.Join(dir, file.Name()), srcType)
		} else {
			if strings.HasSuffix(file.Name(), ".jpg") {
				err := imageconv.Convert(filepath.Join(dir, file.Name()))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}
	}
}
