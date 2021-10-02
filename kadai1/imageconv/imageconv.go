// convert package
// this package convert jpeg to png.

package imageconv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "io"
	"log"
	"os"
	"strings"
)

func Convert(filename string) error {
	img, err := readImage(filename)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return writeImage(img, filename)
}

func readImage(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil { log.Print(err) }
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func writeImage(src image.Image, filepath string) error {
	outfilename := strings.TrimSuffix(filepath, ".jpg") + ".png"
	out, err := os.Create(outfilename)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	png.Encode(out, src)
	return nil
}
