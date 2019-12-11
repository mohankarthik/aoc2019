package computer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputer_Day02_01(t *testing.T) {
	c := NewComputerWithName("Day02_01", []int{1,9,10,3,2,3,11,0,99,30,40,50})
	c.Run()
	assert.Equal(t, []int{3500,9,10,70,2,3,11,0,99,30,40,50}, c.code[0:12])
}

func TestComputer_Day02_02(t *testing.T) {
	c := NewComputerWithName("Day02_02", []int{1,0,0,0,99})
	c.Run()
	assert.Equal(t, []int{2,0,0,0,99}, c.code[0:5])
}

func TestComputer_Day02_03(t *testing.T) {
	c := NewComputerWithName("Day02_03", []int{2,3,0,3,99})
	c.Run()
	assert.Equal(t, []int{2,3,0,6,99}, c.code[0:5])
}

func TestComputer_Day02_04(t *testing.T) {
	c := NewComputerWithName("Day02_04", []int{2,4,4,5,99,0})
	c.Run()
	assert.Equal(t, []int{2,4,4,5,99,9801}, c.code[0:6])
}

func TestComputer_Day02_05(t *testing.T) {
	c := NewComputerWithName("Day02_05", []int{1,1,1,4,99,5,6,0,99})
	c.Run()
	assert.Equal(t, []int{30,1,1,4,2,5,6,0,99}, c.code[0:9])
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

func TestComputerDay0901(t *testing.T) {
	c := NewComputerWithName("test computer", []int{109,2000,109,19,99})
	c.Run()
	assert.Equal(t, 2019, c.relBase)
}

func TestComputerDay0902(t *testing.T) {
	code := []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	c := NewComputerWithName("test computer", code)
	go c.Run()
	cnt := 0
	for {
		res := <- c.Output
		assert.Equal(t, code[cnt], res.Data)
		cnt++
		if cnt == len(code) {
			break
		}
	}
}

func TestComputerDay0903(t *testing.T) {
	c := NewComputerWithName("test computer", []int{1102,34915192,34915192,7,4,7,99,0})
	go c.Run()
	res := <- c.Output
	assert.Equal(t, 1219070632396864, res.Data)
}

func TestComputerDay0904(t *testing.T) {
	c := NewComputerWithName("test computer", []int{104,1125899906842624,99})
	go c.Run()
	res := <- c.Output
	assert.Equal(t, 1125899906842624, res.Data)
}

func TestComputerDay0905(t *testing.T) {
	c := NewComputerWithName("test computer", []int{1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1102,3,1,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1101,0,34,1006,1101,0,689,1022,1102,27,1,1018,1102,1,38,1010,1102,1,31,1012,1101,20,0,1015,1102,1,791,1026,1102,0,1,1020,1101,24,0,1000,1101,0,682,1023,1101,788,0,1027,1101,0,37,1005,1102,21,1,1011,1102,1,28,1002,1101,0,529,1024,1101,39,0,1017,1102,30,1,1013,1101,0,23,1003,1102,524,1,1025,1101,32,0,1007,1102,25,1,1008,1101,29,0,1001,1101,33,0,1016,1101,410,0,1029,1101,419,0,1028,1101,22,0,1014,1102,26,1,1019,1102,1,35,1009,1102,36,1,1004,1102,1,1,1021,109,11,2107,22,-8,63,1005,63,199,4,187,1106,0,203,1001,64,1,64,1002,64,2,64,109,2,21108,40,40,-2,1005,1011,221,4,209,1106,0,225,1001,64,1,64,1002,64,2,64,109,13,21102,41,1,-7,1008,1019,41,63,1005,63,251,4,231,1001,64,1,64,1106,0,251,1002,64,2,64,109,-19,1202,1,1,63,1008,63,26,63,1005,63,271,1105,1,277,4,257,1001,64,1,64,1002,64,2,64,109,7,2101,0,-6,63,1008,63,24,63,1005,63,297,1106,0,303,4,283,1001,64,1,64,1002,64,2,64,109,7,1205,-1,315,1105,1,321,4,309,1001,64,1,64,1002,64,2,64,109,-11,21107,42,41,0,1005,1010,341,1001,64,1,64,1106,0,343,4,327,1002,64,2,64,109,-8,1207,6,24,63,1005,63,363,1001,64,1,64,1106,0,365,4,349,1002,64,2,64,109,11,1206,8,381,1001,64,1,64,1106,0,383,4,371,1002,64,2,64,109,4,1205,4,401,4,389,1001,64,1,64,1105,1,401,1002,64,2,64,109,14,2106,0,-3,4,407,1001,64,1,64,1106,0,419,1002,64,2,64,109,-33,1202,3,1,63,1008,63,29,63,1005,63,445,4,425,1001,64,1,64,1105,1,445,1002,64,2,64,109,-5,2102,1,7,63,1008,63,25,63,1005,63,465,1105,1,471,4,451,1001,64,1,64,1002,64,2,64,109,11,21107,43,44,7,1005,1011,489,4,477,1105,1,493,1001,64,1,64,1002,64,2,64,109,-3,1208,8,35,63,1005,63,511,4,499,1105,1,515,1001,64,1,64,1002,64,2,64,109,25,2105,1,-2,4,521,1106,0,533,1001,64,1,64,1002,64,2,64,109,-8,21108,44,47,-8,1005,1010,549,1106,0,555,4,539,1001,64,1,64,1002,64,2,64,109,-19,1207,7,35,63,1005,63,577,4,561,1001,64,1,64,1106,0,577,1002,64,2,64,109,2,2108,32,0,63,1005,63,597,1001,64,1,64,1106,0,599,4,583,1002,64,2,64,109,13,2101,0,-7,63,1008,63,32,63,1005,63,625,4,605,1001,64,1,64,1105,1,625,1002,64,2,64,109,-13,2107,24,2,63,1005,63,645,1001,64,1,64,1106,0,647,4,631,1002,64,2,64,109,18,21101,45,0,-4,1008,1015,43,63,1005,63,671,1001,64,1,64,1105,1,673,4,653,1002,64,2,64,109,-6,2105,1,10,1001,64,1,64,1105,1,691,4,679,1002,64,2,64,109,1,1208,-6,23,63,1005,63,707,1105,1,713,4,697,1001,64,1,64,1002,64,2,64,109,-2,1206,8,731,4,719,1001,64,1,64,1106,0,731,1002,64,2,64,109,-7,21102,46,1,5,1008,1010,43,63,1005,63,751,1106,0,757,4,737,1001,64,1,64,1002,64,2,64,109,-9,2108,24,4,63,1005,63,779,4,763,1001,64,1,64,1106,0,779,1002,64,2,64,109,38,2106,0,-7,1106,0,797,4,785,1001,64,1,64,1002,64,2,64,109,-27,2102,1,-6,63,1008,63,29,63,1005,63,819,4,803,1105,1,823,1001,64,1,64,1002,64,2,64,109,1,21101,47,0,7,1008,1015,47,63,1005,63,845,4,829,1105,1,849,1001,64,1,64,1002,64,2,64,109,-11,1201,5,0,63,1008,63,31,63,1005,63,869,1106,0,875,4,855,1001,64,1,64,1002,64,2,64,109,5,1201,4,0,63,1008,63,34,63,1005,63,901,4,881,1001,64,1,64,1105,1,901,4,64,99,21102,27,1,1,21101,915,0,0,1105,1,922,21201,1,58905,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21101,0,942,0,1106,0,922,22101,0,1,-1,21201,-2,-3,1,21102,1,957,0,1106,0,922,22201,1,-1,-2,1106,0,968,22102,1,-2,-2,109,-3,2106,0,0})
	go c.Run()
	c.Input <- Msg{Sender: "", Data: 1}
	res := <- c.Output
	assert.Equal(t, 3280416268, res.Data)
}

func TestComputerDay0906(t *testing.T) {
	c := NewComputerWithName("test computer", []int{1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1102,3,1,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1101,0,34,1006,1101,0,689,1022,1102,27,1,1018,1102,1,38,1010,1102,1,31,1012,1101,20,0,1015,1102,1,791,1026,1102,0,1,1020,1101,24,0,1000,1101,0,682,1023,1101,788,0,1027,1101,0,37,1005,1102,21,1,1011,1102,1,28,1002,1101,0,529,1024,1101,39,0,1017,1102,30,1,1013,1101,0,23,1003,1102,524,1,1025,1101,32,0,1007,1102,25,1,1008,1101,29,0,1001,1101,33,0,1016,1101,410,0,1029,1101,419,0,1028,1101,22,0,1014,1102,26,1,1019,1102,1,35,1009,1102,36,1,1004,1102,1,1,1021,109,11,2107,22,-8,63,1005,63,199,4,187,1106,0,203,1001,64,1,64,1002,64,2,64,109,2,21108,40,40,-2,1005,1011,221,4,209,1106,0,225,1001,64,1,64,1002,64,2,64,109,13,21102,41,1,-7,1008,1019,41,63,1005,63,251,4,231,1001,64,1,64,1106,0,251,1002,64,2,64,109,-19,1202,1,1,63,1008,63,26,63,1005,63,271,1105,1,277,4,257,1001,64,1,64,1002,64,2,64,109,7,2101,0,-6,63,1008,63,24,63,1005,63,297,1106,0,303,4,283,1001,64,1,64,1002,64,2,64,109,7,1205,-1,315,1105,1,321,4,309,1001,64,1,64,1002,64,2,64,109,-11,21107,42,41,0,1005,1010,341,1001,64,1,64,1106,0,343,4,327,1002,64,2,64,109,-8,1207,6,24,63,1005,63,363,1001,64,1,64,1106,0,365,4,349,1002,64,2,64,109,11,1206,8,381,1001,64,1,64,1106,0,383,4,371,1002,64,2,64,109,4,1205,4,401,4,389,1001,64,1,64,1105,1,401,1002,64,2,64,109,14,2106,0,-3,4,407,1001,64,1,64,1106,0,419,1002,64,2,64,109,-33,1202,3,1,63,1008,63,29,63,1005,63,445,4,425,1001,64,1,64,1105,1,445,1002,64,2,64,109,-5,2102,1,7,63,1008,63,25,63,1005,63,465,1105,1,471,4,451,1001,64,1,64,1002,64,2,64,109,11,21107,43,44,7,1005,1011,489,4,477,1105,1,493,1001,64,1,64,1002,64,2,64,109,-3,1208,8,35,63,1005,63,511,4,499,1105,1,515,1001,64,1,64,1002,64,2,64,109,25,2105,1,-2,4,521,1106,0,533,1001,64,1,64,1002,64,2,64,109,-8,21108,44,47,-8,1005,1010,549,1106,0,555,4,539,1001,64,1,64,1002,64,2,64,109,-19,1207,7,35,63,1005,63,577,4,561,1001,64,1,64,1106,0,577,1002,64,2,64,109,2,2108,32,0,63,1005,63,597,1001,64,1,64,1106,0,599,4,583,1002,64,2,64,109,13,2101,0,-7,63,1008,63,32,63,1005,63,625,4,605,1001,64,1,64,1105,1,625,1002,64,2,64,109,-13,2107,24,2,63,1005,63,645,1001,64,1,64,1106,0,647,4,631,1002,64,2,64,109,18,21101,45,0,-4,1008,1015,43,63,1005,63,671,1001,64,1,64,1105,1,673,4,653,1002,64,2,64,109,-6,2105,1,10,1001,64,1,64,1105,1,691,4,679,1002,64,2,64,109,1,1208,-6,23,63,1005,63,707,1105,1,713,4,697,1001,64,1,64,1002,64,2,64,109,-2,1206,8,731,4,719,1001,64,1,64,1106,0,731,1002,64,2,64,109,-7,21102,46,1,5,1008,1010,43,63,1005,63,751,1106,0,757,4,737,1001,64,1,64,1002,64,2,64,109,-9,2108,24,4,63,1005,63,779,4,763,1001,64,1,64,1106,0,779,1002,64,2,64,109,38,2106,0,-7,1106,0,797,4,785,1001,64,1,64,1002,64,2,64,109,-27,2102,1,-6,63,1008,63,29,63,1005,63,819,4,803,1105,1,823,1001,64,1,64,1002,64,2,64,109,1,21101,47,0,7,1008,1015,47,63,1005,63,845,4,829,1105,1,849,1001,64,1,64,1002,64,2,64,109,-11,1201,5,0,63,1008,63,31,63,1005,63,869,1106,0,875,4,855,1001,64,1,64,1002,64,2,64,109,5,1201,4,0,63,1008,63,34,63,1005,63,901,4,881,1001,64,1,64,1105,1,901,4,64,99,21102,27,1,1,21101,915,0,0,1105,1,922,21201,1,58905,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21101,0,942,0,1106,0,922,22101,0,1,-1,21201,-2,-3,1,21102,1,957,0,1106,0,922,22201,1,-1,-2,1106,0,968,22102,1,-2,-2,109,-3,2106,0,0})
	go c.Run()
	c.Input <- Msg{Sender: "", Data: 2}
	res := <- c.Output
	fmt.Println(res)
	//assert.Equal(t, 3280416268, res.Data)
}