package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

type Position struct {
	x int
	y int
	z int
}

type Velocity struct {
	x int
	y int
	z int
}

type Moon struct {
	pos Position
	vel Velocity
}

type JupiterSystem struct {
	moons [4]Moon
}

func intAbs(val int) int {
	return int(math.Abs(float64(val)))
}

func computeEnergy(sys JupiterSystem) int {
	energy := 0
	for _,moon := range sys.moons {
		pot := intAbs(moon.pos.x) + intAbs(moon.pos.y) + intAbs(moon.pos.z)
		kin := intAbs(moon.vel.x) + intAbs(moon.vel.y) + intAbs(moon.vel.z)
		energy += pot * kin
	}
	return energy
}

func simulateSysem(sys JupiterSystem) JupiterSystem {
	// Compute the velocities
	for i := range sys.moons {
		for j := range sys.moons {
			if i == j {
				continue
			}

			if sys.moons[i].pos.x < sys.moons[j].pos.x {
				sys.moons[i].vel.x += 1
			} else if sys.moons[i].pos.x > sys.moons[j].pos.x {
				sys.moons[i].vel.x -= 1
			} else {
				
			}

			if sys.moons[i].pos.y < sys.moons[j].pos.y {
				sys.moons[i].vel.y += 1
			} else if sys.moons[i].pos.y > sys.moons[j].pos.y {
				sys.moons[i].vel.y -= 1
			} else {
				
			}

			if sys.moons[i].pos.z < sys.moons[j].pos.z {
				sys.moons[i].vel.z += 1
			} else if sys.moons[i].pos.z > sys.moons[j].pos.z {
				sys.moons[i].vel.z -= 1
			} else {
				
			}
		}
	}

	// Move the moons
	for i := range sys.moons {
		sys.moons[i].pos.x += sys.moons[i].vel.x
		sys.moons[i].pos.y += sys.moons[i].vel.y
		sys.moons[i].pos.z += sys.moons[i].vel.z
	}

	return sys
}

func processInput(data string) (system JupiterSystem) {
	s := strings.Split(data, "\n")
	for i := 0; i < 4; i++ {
		t := strings.Split(s[i][1:len(s[i])-1], ",")
		temp := make([]int, 3)
		for j := 0; j < 3; j++ {
			v := strings.Split(t[j], "=")
			val, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				panic(err)
			}
			temp[j] = int(val)
		}
		system.moons[i].pos.x = temp[0]
		system.moons[i].pos.y = temp[1]
		system.moons[i].pos.z = temp[2]
	}
	return system
}

func main () {
	input := 
`<x=-16, y=15, z=-9>
<x=-14, y=5, z=4>
<x=2, y=0, z=6>
<x=-3, y=18, z=9>`

	sys := processInput(input)
	orig := processInput(input)
	cnt := 0
	for {
		sys = simulateSysem(sys)
		cnt++
		if sys == orig {
			fmt.Println(cnt)
			break
		}
	}
}