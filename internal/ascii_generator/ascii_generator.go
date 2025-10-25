package ascii_generator

import "fmt"

// const brightness_characters string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,`'."
const brightness_characters string = "@GDCIi+-."

func GenerateAsciiArtFromAverages(chunked_average [][][]int) {

	for y_i := 0; y_i < len(chunked_average); y_i++ {
		for x_i := 0; x_i < len(chunked_average[y_i]); x_i++ {
			print_rgb_to_character(chunked_average[y_i][x_i])
		}
		fmt.Println("")
	}
}

func print_rgb_to_character(rgb_value []int) {
	if len(rgb_value) != 3 {
		panic("A rgb value must have exactly 3 entries")
	}
	for i := 0; i < 3; i++ {
		if rgb_value[i] < 0 || rgb_value[i] > 255 {
			panic("Brightness must be between 0 and 255")
		}

	}
	darkness := 255 - (rgb_value[0]+rgb_value[1]+rgb_value[2])/3
	scale_darkness := float32(len(brightness_characters)-1) / float32(255)
	darkness_index := float32(darkness) * (scale_darkness)

	colorReset := "\033[0m"
	color := fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb_value[0], rgb_value[1], rgb_value[2])

	fmt.Print(color, string(brightness_characters[int(darkness_index)]), colorReset)

}
