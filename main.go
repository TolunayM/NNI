package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func loadGrid(filePath string) (grid [][]color.Color) {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}

	return
}

func saveGrid(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}

	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	png.Encode(imgFile, img.SubImage(img.Rect))
}

func resize(grid [][]color.Color, scale float64) (resized [][]color.Color) {

	xlen, ylen := int(float64(len(grid))*scale), int(float64(len(grid[0]))*scale)
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			xp := int(math.Floor(float64(x) / scale))
			yp := int(math.Floor(float64(y) / scale))
			resized[x][y] = grid[xp][yp]
		}
	}
	return
}

func main() {

	grid := loadGrid("pics/WoT1080p.png")
	resized := resize(grid, 0.5)
	saveGrid("pics/WoT720p.png", resized)

}
