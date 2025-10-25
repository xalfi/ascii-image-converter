package main

import (
	"fmt"
	"os"

	_ "image/gif"
	_ "image/png"

	"ascii-image-converter/internal/ascii_generator"
	"ascii-image-converter/internal/image_calculator"
	"ascii-image-converter/internal/image_parser"
)

const width_console int = 120
const width_height_ratio_character float32 = 0.5

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
	chunk_size_x, chunk_size_y := image_calculator.CalculateChunkSizesWithWidth(img, width_console, width_height_ratio_character)
	chunked_brightness := image_calculator.CalculateChunkedBrightness(img, chunk_size_x, chunk_size_y)
	image_parser.SaveChunkedImageAsBlackWhite(chunked_brightness, "output.jpg")
	ascii_generator.GenerateAsciiArtFromBrightness(chunked_brightness)

}
