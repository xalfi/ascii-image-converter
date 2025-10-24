package main

import (
	"fmt"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"ascii-image-converter/internal/image_parser"
)

func main() {
	var args []string = os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path of the image to convert.")
		return
	}
	img := image_parser.ReadImageFromPath(args[1])
	if img == nil {
		return
	}
}
