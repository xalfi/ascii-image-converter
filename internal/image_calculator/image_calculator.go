package image_calculator

import (
	"image"
)

func CalculateChunkSizesWithWidth(img image.Image, new_width int, width_height_ratio float32) (int, int) {
	img_width, img_height := img.Bounds().Max.X, img.Bounds().Max.Y // Assuming that the image always starts at (0,0)
	chunk_size_x := img_width / new_width
	new_height := int(float32(img_height/chunk_size_x) * (width_height_ratio))
	chunk_size_y := img_height / new_height
	return chunk_size_x, chunk_size_y
}

func CalculateChunkedBrightness(img image.Image, chunk_size_x int, chunk_size_y int) [][]int {

	bounds := img.Bounds()

	var chunked_brightness [][]int
	for y_coord := bounds.Min.Y; y_coord < bounds.Max.Y; y_coord += chunk_size_y {
		var brightness_row []int
		for x_coord := bounds.Min.X; x_coord < bounds.Max.X; x_coord += chunk_size_x {

			area := image.Rect(
				x_coord,
				y_coord,
				min(x_coord+chunk_size_x, bounds.Max.X),
				min(y_coord+chunk_size_y, bounds.Max.Y),
			)
			brightness_row = append(brightness_row, calculate_average_brightness(area, img))
		}
		chunked_brightness = append(chunked_brightness, brightness_row)
	}
	return chunked_brightness
}

func calculate_average_brightness(area image.Rectangle, img image.Image) int {
	var total_brightness int64
	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			ri, gi, bi, _ := int(r/257), int(g/257), int(b/257), int(a/257)
			avg_pixel := int((ri + gi + bi) / 3)
			total_brightness += int64(avg_pixel)
		}
	}
	total_avg := total_brightness / int64(area.Dx())
	total_avg = total_avg / int64(area.Dy())

	return int(total_avg)
}
