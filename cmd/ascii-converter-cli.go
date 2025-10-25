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
const DEFAULT_MONOCHROME_MODE bool = false
const DEFAULT_OUTPUT_PATH string = ""

func main() {
	r_flag := flag.Float64("r", DEFAULT_WIDTH_HEIGHT_RATION_CHARACTER, "The width to height ratio of a single character on the console")
	w_flag := flag.Int("w", DEFAULT_WIDTH_CONSOLE, "The number of characters to display on the console in a single row")
	o_flag := flag.String("o", DEFAULT_OUTPUT_PATH, "The path you want to save a jpg of the chunked colors")
	c_flag := flag.Bool("m", DEFAULT_MONOCHROME_MODE, "If you want to use colored output (true or false)")

	flag.Parse()

	width_console := *w_flag
	ratio := *r_flag
	out_path := *o_flag
	do_use_color := !*c_flag

	var args []string = flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide the path of the image to convert as the last argument!")
		return
	}
	img := image_parser.ReadImageFromPath(args[0])
	if img == nil {
		return
	}

	chunk_size_x, chunk_size_y := image_calculator.CalculateChunkSizesWithWidth(img, width_console, float32(ratio))
	chunked_average := image_calculator.CalculateChunkedAverage(img, chunk_size_x, chunk_size_y)
	if len(out_path) > 0 {
		image_parser.SaveChunkedImage(chunked_average, out_path, do_use_color)

	}
	ascii_generator.GenerateAsciiArtFromAverages(chunked_average, do_use_color)

}
