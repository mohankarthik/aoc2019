package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"aoc2019/computer"
	"io/ioutil"
	"strconv"
	"os"
	"strings"
)

func getInput() []int {
	code := make([]int, 0)
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(data), ",")

	for i := range s {
		val, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			panic(err)
		}
		code = append(code, int(val))
	}

	return code
}

type position struct {
	x int
	y int
}

type orientation struct {
	dx int
	dy int
}

func changeOrientation(move int, current orientation) (new orientation) {
	if move == 0 {
		if current.dx == 0 && current.dy == 1 {
			new = orientation{-1, 0}
		} else if current.dx == -1 && current.dy == 0 {
			new = orientation{0, -1}
		} else if current.dx == 0 && current.dy == -1 {
			new = orientation{1, 0}
		} else if current.dx == 1 && current.dy == 0 {
			new = orientation{0, 1}
		} else {
			panic("Unknown orientation")
		}
	} else {
		if current.dx == 0 && current.dy == 1 {
			new = orientation{1, 0}
		} else if current.dx == 1 && current.dy == 0 {
			new = orientation{0, -1}
		} else if current.dx == 0 && current.dy == -1 {
			new = orientation{-1, 0}
		} else if current.dx == -1 && current.dy == 0 {
			new = orientation{0, 1}
		} else {
			panic("Unknown orientation")
		}
	}

	return new
}

func moveRobot(currentPos position, currentOri orientation) (position) {
	return position{currentPos.x + currentOri.dx, currentPos.y + currentOri.dy}
}

func findImageSize(surface map[position]int) (int,int,[][]int) {
	minX := 10000
	minY := 10000
	maxX :=-10000
	maxY :=-10000
	for pos,_ := range surface {
		if minX > pos.x {
			minX = pos.x
		}
		if minY > pos.y {
			minY = pos.y
		}
		if maxX < pos.x {
			maxX = pos.x
		}
		if maxY < pos.y {
			maxY = pos.y
		}
	}
	width := maxX - minX + 1
	height := maxY - minY + 1

	// Make a new surface
	newSurface := make([][]int, height)
	for i := range newSurface {
		newSurface[i] = make([]int, width)
	}

	// Fill the colors
	for pos,col := range surface {
		newSurface[pos.y-minY][pos.x-minX] = col
	}

	// Return the values
	return width, height, newSurface
}

func printRegistration(filename string, pixelSize int, surface map[position]int) {
	width, height, data:= findImageSize(surface)
	
	rect := image.Rect(0, 0, width*pixelSize, height*pixelSize)
	img := image.NewRGBA(rect)
	green := color.RGBA{14, 184, 14, 0xff}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			for y := r * pixelSize; y < (r+1)*pixelSize; y++ {
				for x := c * pixelSize; x < (c+1)*pixelSize; x++ {
					pixel := data[r][c]
					switch pixel {
					case 0:
						img.Set(x, y, color.Black)
					case 1:
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(filename)
	png.Encode(f, img)
}

func runRobot(code []int) map[position]int {
	currentPos := position{0,0}
	currentOri := orientation{0,1}
	surface := make(map[position]int)
	robot := computer.NewComputerWithName("Robot", code)
	surface[currentPos] = 1
	
	go robot.Run()
	for {
		// If we don't have the surface mapped, create it
		if _, ok := surface[currentPos]; !ok {
			surface[currentPos] = 0
		}

		if robot.IsHalted {
			break
		}

		if robot.IsWaitingForInput {
			// Feed input to the robot
			robot.Input <- computer.Msg{Sender:"color", Data:surface[currentPos]}

			// Get the outputs
			color := <- robot.Output
			move := <- robot.Output

			// Paint the surface
			surface[currentPos] = color.Data

			// Move the robot
			currentOri = changeOrientation(move.Data, currentOri)
			currentPos = moveRobot(currentPos, currentOri)
		}
	}

	return surface
}

func main() {
	fmt.Println("Day 11")
	code := getInput()
	surface := runRobot(code)
	fmt.Println(len(surface))
	printRegistration("out.png", 20, surface)
}
