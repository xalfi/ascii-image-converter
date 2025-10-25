# ascii-image-converter

A simple CLI tool to convert images to ASCII art.

## Features

- Convert an image to ASCII art (color and monochrome).
- Save an intermediate chunked image (`output.jpg`).
- Simple CLI with configurable width and character aspect ratio.

## Examples

| Original |Width 60 | Width 120 | Width 250 |
| --- | --- | --- | --- | 
| ![](assets/examples/wednesday/wednesday.jpeg) | ![](assets/examples/wednesday/wednesday_output_w60.jpeg) | ![](assets/examples/wednesday/wednesday_output_w120.jpeg) | ![](assets/examples/wednesday/wednesday_output_w250.jpeg)

## Requirements

- Go (see [go.mod](go.mod) for target version)

## Install

Build the binary:

```sh
go build ./cmd
```

## Usage
Provide a path to an image file as the last argument. Example:
```
# run with defaults
go run ./cmd path/to/image.jpg

# custom width, ratio and monochrome
go run ./cmd -w 80 -r 0.6 -m path/to/image.jpg
```

## CLI flags
- `-w <int>`: The width (num. of character) of the generated ascii image
- `-r <float>`: The ration between width and height of a single character (defaults to 0.5, which is good for most terminal emulators)
- `-m`: Monochrome Mode (generates only black and white output)
- `-o <output.jpg>`: Store the scaled down images to a jpg file (only colors as pixels, no ascii symbols)


## License
This project is licensed under the MIT License — see [LICENSE](LICENSE).


