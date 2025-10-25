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

func CalculateChunkedAverage(img image.Image, chunk_size_x int, chunk_size_y int) [][][]int {

	bounds := img.Bounds()

	var chunked_average [][][]int
	for y_coord := bounds.Min.Y; y_coord < bounds.Max.Y; y_coord += chunk_size_y {
		var chunk_row [][]int
		for x_coord := bounds.Min.X; x_coord < bounds.Max.X; x_coord += chunk_size_x {

			area := image.Rect(
				x_coord,
				y_coord,
				min(x_coord+chunk_size_x, bounds.Max.X),
				min(y_coord+chunk_size_y, bounds.Max.Y),
			)
			chunk_row = append(chunk_row, calculate_average_colors(area, img))
		}
		chunked_average = append(chunked_average, chunk_row)
	}
	return chunked_average
}

// Calculates the average brightness of an area in a given image
func calculate_average_colors(area image.Rectangle, img image.Image) []int {
	var total_r float64 = 0
	var total_g float64 = 0
	var total_b float64 = 0

	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			ri, gi, bi, _ := int(r/257), int(g/257), int(b/257), int(a/257)

			total_r += float64(ri)
			total_g += float64(gi)
			total_b += float64(bi)
		}
	}
	total_r = (total_r / float64(area.Dx())) / float64(area.Dy())
	total_g = (total_g / float64(area.Dx())) / float64(area.Dy())
	total_b = (total_b / float64(area.Dx())) / float64(area.Dy())
	return []int{int(total_r), int(total_g), int(total_b)}
}
