package computer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputer_Day02_01(t *testing.T) {
	c := NewComputerWithName("Day02_01", []int{1,9,10,3,2,3,11,0,99,30,40,50})
	c.Run()
	assert.Equal(t, []int{3500,9,10,70,2,3,11,0,99,30,40,50}, c.code)
}

func TestComputer_Day02_02(t *testing.T) {
	c := NewComputerWithName("Day02_02", []int{1,0,0,0,99})
	c.Run()
	assert.Equal(t, []int{2,0,0,0,99}, c.code)
}

func TestComputer_Day02_03(t *testing.T) {
	c := NewComputerWithName("Day02_03", []int{2,3,0,3,99})
	c.Run()
	assert.Equal(t, []int{2,3,0,6,99}, c.code)
}

func TestComputer_Day02_04(t *testing.T) {
	c := NewComputerWithName("Day02_04", []int{2,4,4,5,99,0})
	c.Run()
	assert.Equal(t, []int{2,4,4,5,99,9801}, c.code)
}

func TestComputer_Day02_05(t *testing.T) {
	c := NewComputerWithName("Day02_05", []int{1,1,1,4,99,5,6,0,99})
	c.Run()
	assert.Equal(t, []int{30,1,1,4,2,5,6,0,99}, c.code)
}

func TestComputer_Day02_06(t *testing.T) {
	c := NewComputerWithName("Day02_06", []int{1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,19,6,23,2,23,13,27,1,27,5,31,2,31,10,35,1,9,35,39,1,39,9,43,2,9,43,47,1,5,47,51,2,13,51,55,1,55,9,59,2,6,59,63,1,63,5,67,1,10,67,71,1,71,10,75,2,75,13,79,2,79,13,83,1,5,83,87,1,87,6,91,2,91,13,95,1,5,95,99,1,99,2,103,1,103,6,0,99,2,14,0,0})
	c.Run()
	assert.Equal(t, 3790645, c.code[0])
}

func TestComputer_Day05_01(t *testing.T) {
	c := NewComputerWithName("Day05_01", []int{3,9,8,9,10,9,4,9,99,-1,8})
	go c.Run()
	c.Input <- Msg{"", 7}
	res := <-c.Output
	assert.Equal(t, 0, res.Data)

	d := NewComputerWithName("Day05_01", []int{3,9,8,9,10,9,4,9,99,-1,8})
	go d.Run()
	d.Input <- Msg{"", 8}
	res = <-d.Output
	assert.Equal(t, 1, res.Data)
}

func TestComputer_Day05_02(t *testing.T) {
	c := NewComputerWithName("Day05_02", []int{3,9,7,9,10,9,4,9,99,-1,8})
	go c.Run()
	c.Input <- Msg{"", 7}
	res := <-c.Output
	assert.Equal(t, 1, res.Data)

	d := NewComputerWithName("Day05_02", []int{3,9,7,9,10,9,4,9,99,-1,8})
	go d.Run()
	d.Input <- Msg{"", 8}
	res = <-d.Output
	assert.Equal(t, 0, res.Data)

	e := NewComputerWithName("Day05_02", []int{3,9,7,9,10,9,4,9,99,-1,8})
	go e.Run()
	e.Input <- Msg{"", -5}
	res = <-e.Output
	assert.Equal(t, 1, res.Data)

	f := NewComputerWithName("Day05_02", []int{3,9,7,9,10,9,4,9,99,-1,8})
	go f.Run()
	f.Input <- Msg{"", 50}
	res = <-f.Output
	assert.Equal(t, 0, res.Data)
}

func TestComputer_Day05_03(t *testing.T) {
	c := NewComputerWithName("Day05_03", []int{3,3,1108,-1,8,3,4,3,99})
	go c.Run()
	c.Input <- Msg{"", 7}
	res := <-c.Output
	assert.Equal(t, 0, res.Data)

	d := NewComputerWithName("Day05_03", []int{3,3,1108,-1,8,3,4,3,99})
	go d.Run()
	d.Input <- Msg{"", 8}
	res = <-d.Output
	assert.Equal(t, 1, res.Data)
}

func TestComputer_Day05_04(t *testing.T) {
	c := NewComputerWithName("Day05_04", []int{3,3,1107,-1,8,3,4,3,99})
	go c.Run()
	c.Input <- Msg{"", 7}
	res := <-c.Output
	assert.Equal(t, 1, res.Data)

	d := NewComputerWithName("Day05_04", []int{3,3,1107,-1,8,3,4,3,99})
	go d.Run()
	d.Input <- Msg{"", 8}
	res = <-d.Output
	assert.Equal(t, 0, res.Data)

	e := NewComputerWithName("Day05_04", []int{3,3,1107,-1,8,3,4,3,99})
	go e.Run()
	e.Input <- Msg{"", -5}
	res = <-e.Output
	assert.Equal(t, 1, res.Data)

	f := NewComputerWithName("Day05_04", []int{3,3,1107,-1,8,3,4,3,99})
	go f.Run()
	f.Input <- Msg{"", 50}
	res = <-f.Output
	assert.Equal(t, 0, res.Data)
}

func TestComputer_Day05_05(t *testing.T) {
	c := NewComputerWithName("Day05_05", []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9})
	go c.Run()
	c.Input <- Msg{"", 0}
	res := <-c.Output
	assert.Equal(t, 0, res.Data)

	d := NewComputerWithName("Day05_05", []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9})
	go d.Run()
	d.Input <- Msg{"", -10}
	res = <-d.Output
	assert.Equal(t, 1, res.Data)
}

func TestComputer_Day05_06(t *testing.T) {
	c := NewComputerWithName("Day05_06", []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1})
	go c.Run()
	c.Input <- Msg{"", 0}
	res := <-c.Output
	assert.Equal(t, 0, res.Data)

	d := NewComputerWithName("Day05_06", []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1})
	go d.Run()
	d.Input <- Msg{"", 200}
	res = <-d.Output
	assert.Equal(t, 1, res.Data)
}

func TestComputer_Day05_07(t *testing.T) {
	c := NewComputerWithName("Day05_07", []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99})
	go c.Run()
	c.Input <- Msg{"", 7}
	res := <-c.Output
	assert.Equal(t, 999, res.Data)

	d := NewComputerWithName("Day05_07", []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99})
	go d.Run()
	d.Input <- Msg{"", 8}
	res = <-d.Output
	assert.Equal(t, 1000, res.Data)

	e := NewComputerWithName("Day05_07", []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99})
	go e.Run()
	e.Input <- Msg{"", 9}
	res = <-e.Output
	assert.Equal(t, 1001, res.Data)
}

func TestComputer_Day05_08(t *testing.T) {
	d := NewComputerWithName("Day05_08", []int{3,225,1,225,6,6,1100,1,238,225,104,0,1102,46,47,225,2,122,130,224,101,-1998,224,224,4,224,1002,223,8,223,1001,224,6,224,1,224,223,223,1102,61,51,225,102,32,92,224,101,-800,224,224,4,224,1002,223,8,223,1001,224,1,224,1,223,224,223,1101,61,64,225,1001,118,25,224,101,-106,224,224,4,224,1002,223,8,223,101,1,224,224,1,224,223,223,1102,33,25,225,1102,73,67,224,101,-4891,224,224,4,224,1002,223,8,223,1001,224,4,224,1,224,223,223,1101,14,81,225,1102,17,74,225,1102,52,67,225,1101,94,27,225,101,71,39,224,101,-132,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1002,14,38,224,101,-1786,224,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1,65,126,224,1001,224,-128,224,4,224,1002,223,8,223,101,6,224,224,1,224,223,223,1101,81,40,224,1001,224,-121,224,4,224,102,8,223,223,101,4,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1008,677,226,224,1002,223,2,223,1005,224,329,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,344,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,359,1001,223,1,223,1108,226,226,224,1002,223,2,223,1006,224,374,101,1,223,223,107,226,226,224,1002,223,2,223,1005,224,389,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,404,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,1107,677,226,224,102,2,223,223,1005,224,434,1001,223,1,223,108,226,677,224,102,2,223,223,1006,224,449,1001,223,1,223,8,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1007,677,226,224,1002,223,2,223,1006,224,479,1001,223,1,223,1007,677,677,224,1002,223,2,223,1005,224,494,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,509,101,1,223,223,1108,226,677,224,102,2,223,223,1005,224,524,1001,223,1,223,7,226,226,224,102,2,223,223,1005,224,539,1001,223,1,223,8,677,677,224,1002,223,2,223,1005,224,554,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,569,1001,223,1,223,7,226,677,224,1002,223,2,223,1005,224,584,1001,223,1,223,1008,226,226,224,1002,223,2,223,1006,224,599,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,629,1001,223,1,223,8,226,677,224,1002,223,2,223,1006,224,644,101,1,223,223,1007,226,226,224,102,2,223,223,1005,224,659,101,1,223,223,108,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226})
	go d.Run()
	d.Input <- Msg{"", 5}
	res := <-d.Output
	assert.Equal(t, 7704130, res.Data)
}

func TestComputer_Run(t *testing.T) {
	c := NewComputerWithName("test computer", []int{3, 0, 4, 0, 99})
	c.Output = nil
	go func() {
		c.Input <- Msg{"", 2}
	}()

	c.Run()
	assert.Equal(t, 2, c.GetLastOutput())
}

func TestComputer_Run4(t *testing.T) {
	c := NewComputerWithName("test computer", []int{101, 1, 6, 7, 4, 7, 99, 0})
	c.Output = make(chan Msg)
	go c.Run()
	res := <-c.Output
	assert.Equal(t, 100, res.Data)
}

func TestComputer_Run3(t *testing.T) {
	c := NewComputerWithName("test computer", []int{1002, 4, 2, 5, 99, 0})

	c.Run()
	assert.Equal(t, []int{1002, 4, 2, 5, 99, 198}, c.code)
}
