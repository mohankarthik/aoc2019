package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

const (
	imageWidth  = 25
	imageHeight = 6
)

const (
	colorBlack = 0
	colorWhite = 1
	colorTrans = 2
)

func getLayers(image string, width int, height int) []string {
	if len(image)%(width*height) != 0 {
		panic("Image size is incorrect")
	}

	layers := make([]string, 0)
	idx := (width * height)
	prevIdx := 0
	for {
		layers = append(layers, image[prevIdx:idx])
		prevIdx = idx
		idx += (width * height)
		if idx > len(image) {
			break
		}
	}

	return layers
}

func countDigits(layers []string) [][]int {
	cnt := make([][]int, 3)

	for i := range layers {
		cnt[0] = append(cnt[0], strings.Count(layers[i], "0"))
		cnt[1] = append(cnt[1], strings.Count(layers[i], "1"))
		cnt[2] = append(cnt[2], strings.Count(layers[i], "2"))
	}

	return cnt
}

func findLayerWithLeastZero(counts [][]int) int {
	minVal := 100000
	minLayer := -1
	for i := range counts[0] {
		if counts[0][i] < minVal {
			minVal = counts[0][i]
			minLayer = i
		}
	}

	return minLayer
}

// Step 1
func checkImage(image string, width int, height int) int {
	layers := getLayers(image, width, height)
	cnt := countDigits(layers)
	idx := findLayerWithLeastZero(cnt)

	return cnt[1][idx] * cnt[2][idx]
}

// Step 2
func getColors(layers []string, width int, height int) string {
	var merge string

	for pix := 0; pix < (width * height); pix++ {
		for layer := 0; layer < len(layers); layer++ {
			if (layers[layer][pix] != colorTrans+48) || (layer == len(layers)-1) {
				merge = merge + string(layers[layer][pix])
				break
			}
		}
	}

	return merge
}

func printLayers(data string, width int, height int) {
	layers := getLayers(data, width, height)

	for i := 0; i < len(layers); i++ {
		fmt.Println("Layer ", i)

		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if layers[i][(row*width)+col] == 49 {
					fmt.Print("\u2588")
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func renderImage(data string, width int, height int, pixelSize int, filename string) {
	rect := image.Rect(0, 0, width*pixelSize, height*pixelSize)
	img := image.NewRGBA(rect)
	// blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{14, 184, 14, 0xff}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			for y := r * pixelSize; y < (r+1)*pixelSize; y++ {
				for x := c * pixelSize; x < (c+1)*pixelSize; x++ {
					pixel := data[r*width+c]
					switch pixel {
					case colorBlack + 48:
						img.Set(x, y, color.Black)
					case colorWhite + 48:
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(filename)
	png.Encode(f, img)
}

func getInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	data := getInput("input.txt")
	layers := getLayers(data, imageWidth, imageHeight)
	printLayers(data, imageWidth, imageHeight)

	checksum := checkImage(data, imageWidth, imageHeight)
	merge := getColors(layers, imageWidth, imageHeight)
	fmt.Println(checksum, merge)
	printLayers(merge, imageWidth, imageHeight)
	renderImage(merge, imageWidth, imageHeight, 20, "out.png")
}
