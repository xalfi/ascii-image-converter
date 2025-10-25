package ascii_generator

import "fmt"

// const brightness_characters string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,`'."
const brightness_characters string = "@#0CIi-. "

func GenerateAsciiArtFromBrightness(chunked_brightness [][]int) {

	for y_i := 0; y_i < len(chunked_brightness); y_i++ {
		for x_i := 0; x_i < len(chunked_brightness[y_i]); x_i++ {
			fmt.Print(brightness_to_character(chunked_brightness[y_i][x_i]))
		}
		fmt.Println("")
	}
}

func brightness_to_character(brightness int) string {
	if brightness < 0 || brightness > 255 {
		panic("Brightness must be between 0 and 255")
	}
	darkness := 255 - brightness
	scale_darkness := float32(len(brightness_characters)-1) / float32(255)
	darkness_index := float32(darkness) * (scale_darkness)

	return string(brightness_characters[int(darkness_index)])

}
