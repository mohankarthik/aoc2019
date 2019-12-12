package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1001(t *testing.T) {
	fmt.Println("Tests starting")
	asteroids := inputToAsteroid(
		`.#..#
		.....
		#####
		....#
		...##`,
	)
	idx, numLos := checkLos(asteroids)
	assert.Equal(t, position{3, 4}, asteroids[idx])
	assert.Equal(t, 8, numLos)
}

func TestDay1002(t *testing.T) {
	asteroids := inputToAsteroid(
		`......#.#.
		#..#.#....
		..#######.
		.#.#.###..
		.#..#.....
		..#....#.#
		#..#....#.
		.##.#..###
		##...#..#.
		.#....####`,
	)
	idx, numLos := checkLos(asteroids)
	assert.Equal(t, position{5, 8}, asteroids[idx])
	assert.Equal(t, 33, numLos)
}

func TestDay1003(t *testing.T) {
	asteroids := inputToAsteroid(
		`#.#...#.#.
		.###....#.
		.#....#...
		##.#.#.#.#
		....#.#.#.
		.##..###.#
		..#...##..
		..##....##
		......#...
		.####.###.`,
	)
	idx, numLos := checkLos(asteroids)
	assert.Equal(t, position{1, 2}, asteroids[idx])
	assert.Equal(t, 35, numLos)
}

func TestDay1004(t *testing.T) {
	asteroids := inputToAsteroid(
		`.#..#..###
		####.###.#
		....###.#.
		..###.##.#
		##.##.#.#.
		....###..#
		..#.#..#.#
		#..#.#.###
		.##...##.#
		.....#.#..`,
	)
	idx, numLos := checkLos(asteroids)
	assert.Equal(t, position{6, 3}, asteroids[idx])
	assert.Equal(t, 41, numLos)
}

func TestDay1005(t *testing.T) {
	asteroids := inputToAsteroid(
		`.#..##.###...#######
		##.############..##.
		.#.######.########.#
		.###.#######.####.#.
		#####.##.#.##.###.##
		..#####..#.#########
		####################
		#.####....###.#.#.##
		##.#################
		#####.##.###..####..
		..######..##.#######
		####.##.####...##..#
		.#####..#.######.###
		##...#.##########...
		#.##########.#######
		.####.#.###.###.#.##
		....##.##.###..#####
		.#.#.###########.###
		#.#.#.#####.####.###
		###.##.####.##.#..##`,
	)
	idx, numLos := checkLos(asteroids)
	assert.Equal(t, position{11, 13}, asteroids[idx])
	assert.Equal(t, 210, numLos)
}

func TestDay10Angle(t *testing.T) {
	assert.Equal(t, getAngle(position{8,3}, position{8,2}), -0.0)
	assert.Equal(t, getAngle(position{8,3}, position{9,2}), 0.7853981633974483)
	assert.Equal(t, getAngle(position{8,3}, position{9,3}), 1.5707963267948966)
	assert.Equal(t, getAngle(position{8,3}, position{9,4}), 2.356194490192345)
	assert.Equal(t, getAngle(position{8,3}, position{8,4}), 3.141592653589793)
	assert.Equal(t, getAngle(position{8,3}, position{7,4}), 3.9269908169872414)
	assert.Equal(t, getAngle(position{8,3}, position{7,3}), 4.71238898038469)
	assert.Equal(t, getAngle(position{8,3}, position{7,2}), 5.497787143782138)
}

func TestDay1006(t *testing.T) {
	asteroids := inputToAsteroid(
		`.#....#####...#..
		##...##.#####..##
		##...#...#.#####.
		..#.....X...###..
		..#.#.....#....##`,
	)
	sequence := sweepAndDestroy(asteroids, position{8,3})
	assert.Equal(t, sequence[0], position{8,1})
	assert.Equal(t, sequence[1], position{9,0})
	assert.Equal(t, sequence[2], position{9,1})
	assert.Equal(t, sequence[3], position{10,0})
	assert.Equal(t, sequence[4], position{9,2})
	assert.Equal(t, sequence[5], position{11,1})
	assert.Equal(t, sequence[6], position{12,1})
	assert.Equal(t, sequence[7], position{11,2})
	assert.Equal(t, sequence[8], position{15,1})
}

func TestDay1007(t *testing.T) {
	asteroids := inputToAsteroid(
		`.#..##.###...#######
		##.############..##.
		.#.######.########.#
		.###.#######.####.#.
		#####.##.#.##.###.##
		..#####..#.#########
		####################
		#.####....###.#.#.##
		##.#################
		#####.##.###..####..
		..######..##.#######
		####.##.####...##..#
		.#####..#.######.###
		##...#.##########...
		#.##########.#######
		.####.#.###.###.#.##
		....##.##.###..#####
		.#.#.###########.###
		#.#.#.#####.####.###
		###.##.####.##.#..##`,
	)
	idx, _ := checkLos(asteroids)
	sweepAndDestroy(asteroids, asteroids[idx])
	//assert.Equal(t, position{11, 12}, sequence[0])
	//assert.Equal(t, position{12, 1}, sequence[1])
	//assert.Equal(t, position{12, 2}, sequence[2])
}