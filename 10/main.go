package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type position struct {
	x int
	y int
}

func inputToAsteroid(data string) []position {
	asteroids := make([]position, 0)
 
	row := 0
	col := 0
	for i := range data {
		if data[i] == 13 {

		} else if data[i] == 10 {
			row++
			col = 0
		} else if data[i] == 35 {
			asteroids = append(asteroids, position{x: col, y: row})
			col++
		} else if data[i] == 46 {
			col++
		}
	}
	return asteroids
}

func getInput() []byte {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return d
}

func getDistance(a position, b position) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func getAngle(a position, b position) float64 {
	angle := math.Atan2(float64(a.x - b.x), float64(a.y - b.y)) * -1
	if (angle < 0) {
		angle *= -1
		angle = (2 * math.Pi) - angle
	}
	return angle
}

func checkLos(asteroids []position) (int, int) {
	numLos := make([]int, len(asteroids))
	max := 0
	maxIdx := -1	

	for i := range asteroids {
		slopes := make([]float64, 0)
		numLos[i] = 0

		for j := range asteroids {
			if j == i {
				continue
			}

			d := getAngle(asteroids[i], asteroids[j])
			bPossible := true
			for _,s := range slopes {
				if d == s {
					bPossible = false
				}
			}
			if bPossible {
				slopes = append(slopes, d)
				numLos[i]++
			}
		}

		if (numLos[i] > max) {
			max = numLos[i]
			maxIdx = i
		}
	}

	return maxIdx, max
}

func sweepAndDestroy(asteroids []position, origin position) (destroyed []position) {
	targets := make(map[float64][]position)
	for _, a := range asteroids {
		if a.x == origin.x && a.y == origin.y {
			continue
		}

		// Calculate the angle wrt origin, offsetting by 45 degrees so that directly upwards counts as 0. Also, since radians go counter-clockwise, multiply by -1
		angle := getAngle(origin, a)

		// Create the map if it's not already there
		if _, ok := targets[angle]; !ok {
			targets[angle] = make([]position, 0)
		}

		// Store the target
		targets[angle] = append(targets[angle], a)

		// Sort ascending from closest to origin
		sort.Slice(targets[angle], func(i, j int) bool {
			return getDistance(origin, targets[angle][i]) < getDistance(origin, targets[angle][j])
		})
	}

	// Order the angles ascending
	order := make([]float64, 0)
	for d := range targets {
		order = append(order, d)
	}
	sort.Float64s(order)

	// Kill 'em all
	i := 0
	for len(destroyed) < len(asteroids)-1 {
		d := order[i%len(order)]
		if len(targets[d]) > 0 {
			destroyed = append(destroyed, targets[d][0])
			targets[d] = targets[d][1:]
		}
		i++
	}

	return destroyed
}

func main() {
	data := getInput()
	asteroids := inputToAsteroid(string(data))
	idx, _ := checkLos(asteroids)
	sequence := sweepAndDestroy(asteroids, asteroids[idx])
	fmt.Println(sequence[199].x*100 + sequence[199].y)
}