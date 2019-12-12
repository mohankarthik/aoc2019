package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1201 (t *testing.T) {
	fmt.Println("Test started")
	input := 
`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

	sys := processInput(input)
	for i := 0; i < 10; i++ {
		sys = simulateSysem(sys)
	}
	energy := computeEnergy(sys)
	assert.Equal(t, energy, 179)
}

func TestDay1202 (t *testing.T) {
	input := 
`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`

	sys := processInput(input)
	for i := 0; i < 100; i++ {
		sys = simulateSysem(sys)
	}
	energy := computeEnergy(sys)
	assert.Equal(t, energy, 1940)
}

func TestDay1203 (t *testing.T) {
	input := 
`<x=-16, y=15, z=-9>
<x=-14, y=5, z=4>
<x=2, y=0, z=6>
<x=-3, y=18, z=9>`

	sys := processInput(input)
	for i := 0; i < 1000; i++ {
		sys = simulateSysem(sys)
	}
	energy := computeEnergy(sys)
	assert.Equal(t, energy, 10664)
}

func TestDay1204 (t *testing.T) {
	input := 
`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

	sys := processInput(input)
	orig := processInput(input)
	cnt := 0
	for {
		sys = simulateSysem(sys)
		cnt++
		if sys == orig {
			break
		}
		if cnt > 2773 {
			panic("err")
		}
	}
	assert.Equal(t, cnt, 2772)
}