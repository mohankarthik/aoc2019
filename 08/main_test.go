package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain01(t *testing.T) {
	got := checkImage("123456789012", 3, 2)
	assert.Equal(t, 1, got)
}

func TestMain02(t *testing.T) {
	got := checkImage("120406219012", 3, 2)
	assert.Equal(t, 4, got)
}

func TestMain03(t *testing.T) {
	data := getInput("input.txt")
	got := checkImage(data, imageWidth, imageHeight)
	assert.Equal(t, 1215, got)
}

func TestMain04(t *testing.T) {
	data := "0222112222120000"
	layers := getLayers(data, 2, 2)
	got := getColors(layers, 2, 2)
	assert.Equal(t, "0110", got)
}

func TestMain05(t *testing.T) {
	data := "0222112222120002"
	layers := getLayers(data, 2, 2)
	got := getColors(layers, 2, 2)
	assert.Equal(t, "0112", got)
}
