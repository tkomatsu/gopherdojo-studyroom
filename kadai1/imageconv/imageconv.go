// convert package
// this package convert jpeg to png.

package imageconv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Convert(dir string, srcType string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", dir, err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			fmt.Println()
			Convert(filepath.Join(dir, file.Name()), srcType)
		} else {
			if strings.HasSuffix(file.Name(), ".jpg") {
				jpg2png(filepath.Join(dir, file.Name()))
			}
		}
	}
}

func jpg2png(filename string) {
	file, err := os.Open(filename)
	if err != nil { log.Print(err) }
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil { log.Print(err) }
	fmt.Fprintf(os.Stderr, "format: %v", format)

	outfilename := strings.TrimSuffix(filename, ".jpg") + ".png"

	out, err := os.Create(outfilename)

	png.Encode(out, img)
}
