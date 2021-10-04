// convert package
// this package convert jpeg to png.

package imageconv

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func Convert(filename, dest, src string) error {
	dest = strings.ToLower(dest)
	src = strings.ToLower(src)

	switch src {
	case "jpeg", "jpg", "png":
	default:
		return errors.New("Invalid format")
	}

	switch dest {
	case "jpeg", "jpg", "png":
	default:
		return errors.New("Invalid format")
	}

	if !strings.HasSuffix(filename, "." + src) {
		return nil
	}

	img, err := readImage(filename, src)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return writeImage(img, filename, dest, src)
}

func readImage(filepath, src string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil { log.Print(err) }
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func writeImage(img image.Image, filepath, dest, src string) error {
	outfilename := strings.TrimSuffix(filepath, "." + src) + "." + dest
	out, err := os.Create(outfilename)
	if err != nil { return err }
	defer out.Close()

	switch dest {
	case "png":
		png.Encode(out, img)
	case "jpg", "jpeg":
		jpeg.Encode(out, img, nil)
	}
	return nil
}
