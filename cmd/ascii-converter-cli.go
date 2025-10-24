package main

import (
	"fmt"
	"image"
	"os"

	"image/color"
	"image/jpeg"

	_ "image/gif"
	_ "image/png"

	"ascii-image-converter/internal/image_calculator"
	"ascii-image-converter/internal/image_parser"
)

const chunk_size_x = 10
const chunk_size_y = 10

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
	fmt.Print("Generated Chunked Brightness: ")
	fmt.Print("Height: ", len(chunked_brightness))
	fmt.Print("Width: ", len(chunked_brightness[0]))

	new_image := image.NewGray(image.Rect(0, 0, len(chunked_brightness[0]), len(chunked_brightness)))

	for y_coord := 0; y_coord < len(chunked_brightness); y_coord++ {
		for x_coord := 0; x_coord < len(chunked_brightness[0]); x_coord++ {
			new_image.Set(x_coord, y_coord, color.Gray{uint8(chunked_brightness[y_coord][x_coord])})
		}
	}
	outFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	error_encode := jpeg.Encode(outFile, new_image, nil)
	if error_encode != nil {
		panic(error_encode)
	}

}
