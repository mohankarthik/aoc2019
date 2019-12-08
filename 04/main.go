package main

import (
	"fmt"
)

func check(number int) bool {
	digits := make([]int, 6)
	dict := make(map[int]int)

	for i := 0; i < 6; i++ {
		digits[i] = number % 10
		number /= 10
	}

	for i := 0; i < 5; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
		dict[digits[i]]++
	}
	dict[digits[len(digits)-1]]++

	for _, v := range dict {
		if v == 2 {
			return true
		}
	}
	return false
}

func main() {
	inputMin := 158126
	inputMax := 624574

	//fmt.Println(check(222555))

	cnt := 0
	for i := inputMin; i <= inputMax; i++ {
		if check(i) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
