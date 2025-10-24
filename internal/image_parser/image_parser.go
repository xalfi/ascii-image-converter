package image_parser

import (
	"fmt"
	"image"
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
