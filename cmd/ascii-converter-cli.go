package main

import (
	"flag"
	"fmt"

	"ascii-image-converter/internal/ascii_generator"
	"ascii-image-converter/internal/image_calculator"
	"ascii-image-converter/internal/image_parser"
)

const DEFAULT_WIDTH_CONSOLE int = 120
const DEFAULT_WIDTH_HEIGHT_RATION_CHARACTER float64 = 0.5
const DEFAULT_COLOR_Mode bool = true

func main() {
	r_flag := flag.Float64("r", DEFAULT_WIDTH_HEIGHT_RATION_CHARACTER, "The width to height ratio of a single character on the console")
	w_flag := flag.Int("w", DEFAULT_WIDTH_CONSOLE, "The number of characters to display on the console in a single row")
	c_flag := flag.Bool("c", DEFAULT_COLOR_Mode, "If you want to use colored output (true or false)")
	flag.Parse()

	width_console := *w_flag
	ratio := *r_flag
	_ = *c_flag

	var args []string = flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide the path of the image to convert.")
		return
	}
	img := image_parser.ReadImageFromPath(args[0])
	if img == nil {
		return
	}
	chunk_size_x, chunk_size_y := image_calculator.CalculateChunkSizesWithWidth(img, width_console, float32(ratio))
	chunked_average := image_calculator.CalculateChunkedAverage(img, chunk_size_x, chunk_size_y)
	image_parser.SaveChunkedImage(chunked_average, "output.jpg")
	ascii_generator.GenerateAsciiArtFromAverages(chunked_average)

	//fmt.Println(string(colorRed), "test", string(colorReset))
	//fmt.Println("next")

}
