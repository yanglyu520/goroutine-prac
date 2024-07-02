package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" //register PNG decoder
	"io"
	"os"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func main() {
	outputFormat := flag.String("format")
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg:%w\n", err)
		os.Exit(1)
	}
}
