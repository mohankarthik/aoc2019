package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain01(t *testing.T) {
	code := []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}
	res := runAmps(code, [5]int{4,3,2,1,0})
	assert.Equal(t, 43210, res)
}

func TestMain02(t *testing.T) {
	code := []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}
	res := runAmps(code, [5]int{0,1,2,3,4})
	assert.Equal(t, 54321, res)
}

func TestMain03(t *testing.T) {
	code := []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}
	res := runAmps(code, [5]int{1,0,4,3,2})
	assert.Equal(t, 65210, res)
}