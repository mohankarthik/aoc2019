package main

import (
	computer "aoc2019/computer"
	"fmt"
	"strconv"
	"io/ioutil"
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
		val, err := strconv.ParseInt(s[i], 10, 32)
		if err != nil {
			panic(err)
		}
		code = append(code, int(val))
	}

	return code
}

func runAmps(code []int, phase [5]int) int {
	amps := make([]computer.Computer, 5)
	
	for i := 0; i < 5; i++ {
		amps[i] = computer.NewComputerWithName("day07 amp " + strconv.Itoa(i), code)
		go amps[i].Run()
		amps[i].Input <- computer.Msg{"phase", phase[i]}
	}

	for i := 0; i < 5; i++ {
		if i == 0 {
			amps[i].Input <- computer.Msg{"init", 0}
		} else {
			res := <- amps[i-1].Output
			amps[i].Input <- res
		}
	}
	res := <- amps[4].Output
	return res.Data
}

func genPhaseCombination(out chan [5]int) {
	var msg [5]int

	for a := 0; a < 5; a++ {
		msg[0] = a
		for b := 0; b < 5; b++ {
			msg[1] = b
			for c := 0; c < 5; c++ {
				msg[2] = c
				for d := 0; d < 5; d++ {
					msg[3] = d
					for e := 0; e < 5; e++ {
						msg[4] = e
						out <- msg
					}
				}
			}
		}
	}
}

func main() {
	code := getInput()
	phaseMsg := make (chan [5]int)
	go genPhaseCombination(phaseMsg)
	maxVal := 0
	for {
		phase := <- phaseMsg
		val := runAmps(code, phase)
		if (val > maxVal) {
			maxVal = val
			fmt.Println(phase, ": ", val)
		}
	}
}
