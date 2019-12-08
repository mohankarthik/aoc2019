package main

import (
	computer "aoc2019/computer"
	"fmt"
	"io/ioutil"
	"strconv"
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

func main() {
	code := getInput()
	comp := computer.NewComputerWithName("day07", code)
	go comp.Run()
	comp.Input <- computer.Msg{Sender: "init", Data: 5}
	res := <-comp.Output
	fmt.Println(res.Data)
}
