package image_parser

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
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

func SaveChunkedImage(chunked_average [][][]int, out_filename string) {

	new_image := image.NewRGBA(image.Rect(0, 0, len(chunked_average[0]), len(chunked_average)))

	for y_coord := 0; y_coord < len(chunked_average); y_coord++ {
		for x_coord := 0; x_coord < len(chunked_average[0]); x_coord++ {
			new_image.Set(x_coord, y_coord,
				color.RGBA{
					uint8(chunked_average[y_coord][x_coord][0]),
					uint8(chunked_average[y_coord][x_coord][1]),
					uint8(chunked_average[y_coord][x_coord][2]),
					0,
				},
			)
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
