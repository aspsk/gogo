package main

import (
	"fmt"
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outputKind = flag.String("kind", "jpeg", "The output kind")

func convert(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	switch *outputKind {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported output format: %s", *outputKind)
	}

	panic("no way")
	return nil
}

func main() {
	flag.Parse()

	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
}
