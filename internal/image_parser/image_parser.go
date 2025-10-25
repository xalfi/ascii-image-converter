package image_parser

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func ReadImageFromPath(path string) image.Image {
	fmt.Printf("Reading image from path: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Could not open the given file", err.Error())
		return nil
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Could not parse the given file:", err.Error())
		return nil
	}
	return img

}

func SaveChunkedImageAsBlackWhite(chunked_brightness [][]int, out_filename string) {

	new_image := image.NewGray(image.Rect(0, 0, len(chunked_brightness[0]), len(chunked_brightness)))

	for y_coord := 0; y_coord < len(chunked_brightness); y_coord++ {
		for x_coord := 0; x_coord < len(chunked_brightness[0]); x_coord++ {
			new_image.Set(x_coord, y_coord, color.Gray{uint8(chunked_brightness[y_coord][x_coord])})
		}
	}
	outFile, err := os.Create(out_filename)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	error_encode := jpeg.Encode(outFile, new_image, nil)
	if error_encode != nil {
		panic(error_encode)
	}

}
