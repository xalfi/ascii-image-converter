package main

import (
	"fmt"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"ascii-image-converter/internal/image_calculator"
	"ascii-image-converter/internal/image_parser"
)

const chunk_size_x = 20
const chunk_size_y = 20

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
	chunked_brightness := image_calculator.CalculateChunkedBrightness(img, chunk_size_x, chunk_size_y)
	fmt.Println(chunked_brightness)

}
