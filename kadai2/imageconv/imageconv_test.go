package imageconv_test

import (
	"testing"

	"github.com/tkomatsu/gopherdojo-studyroom/kadai2/imageconv"
)

var testCase = []struct {
	name     string
	file     string
	in       string
	out      string
	expected error
}{
	{name: "png to jpg", file: "../test_images/image.png", in: "png", out: "jpg", expected: nil},
	{name: "jpg to png", file: "../test_images/image.jpg", in: "jpg", out: "png", expected: nil},
	{name: "jpeg to png", file: "../test_images/image.jpeg", in: "jpeg", out: "png", expected: nil},
	{name: "png to jpeg", file: "../test_images/image.png", in: "png", out: "jpeg", expected: nil},
}

func TestImageconv(t *testing.T) {
	t.Helper()
	for _, c := range testCase {
		t.Run(c.name, func(t *testing.T) {
			err := imageconv.Convert(c.file, c.in, c.out)
			if err != c.expected {
				t.Errorf("case: %v\n", c)
			}
		})
	}
}
