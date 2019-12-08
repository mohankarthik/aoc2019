package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Parameter modes
const (
	ParamModePosition  = 0
	ParamModeImmediate = 1
)

// OP Codes of the IntCode Computer
const (
	OpAdd       = 1
	OpMul       = 2
	OpIn        = 3
	OpOut       = 4
	OpJumpTrue  = 5
	OpJumpFalse = 6
	OpLessThan  = 7
	OpEquals    = 8
	OpHalt      = 99
)

func decodeOp(opCode int) []int {
	result := make([]int, 4)

	// Get the actual opCode
	val := opCode % 100
	result[0] = val
	opCode /= 100

	for i := 1; i < 10; i++ {
		if opCode <= 0 {
			break
		}
		val = opCode % 10
		result[i] = val
		opCode /= 10
	}

	return result
}

func getOperand(IntCode []int, pos int, mode int) int {
	if mode == ParamModePosition {
		return IntCode[IntCode[pos]]
	}
	return IntCode[pos]
}

func runComputer(IntCode []int) []int {
	pc := 0

	// Go into an infinite loop
	for {
		// Get the Op Code
		opCode := decodeOp(IntCode[pc])

		// Check for halt
		if opCode[0] == OpHalt {
			return IntCode
		}

		if opCode[0] == OpAdd {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])
			op3 := IntCode[pc+3]

			// Execute
			IntCode[op3] = op1 + op2

			// Increment the program counter
			pc += 4
		} else if opCode[0] == OpMul {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])
			op3 := IntCode[pc+3]

			// Execute
			IntCode[op3] = op1 * op2

			// Increment
			pc += 4
		} else if opCode[0] == OpIn {
			// Execute
			// TODO Fix this
			IntCode[IntCode[pc+1]] = 5

			// Increment
			pc += 2
		} else if opCode[0] == OpOut {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])

			// Execute
			fmt.Println(op1)

			// Increment
			pc += 2
		} else if opCode[0] == OpJumpTrue {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])

			// Execute
			if op1 != 0 {
				pc = op2
			} else {
				// Increment
				pc += 3
			}
		} else if opCode[0] == OpJumpFalse {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])

			// Execute
			if op1 == 0 {
				pc = op2
			} else {
				// Increment
				pc += 3
			}
		} else if opCode[0] == OpLessThan {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])
			op3 := IntCode[pc+3]

			// Execute
			if op1 < op2 {
				IntCode[op3] = 1
			} else {
				IntCode[op3] = 0
			}

			// Increment
			pc += 4
		} else if opCode[0] == OpEquals {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])
			op2 := getOperand(IntCode, pc+2, opCode[2])
			op3 := IntCode[pc+3]

			// Execute
			if op1 == op2 {
				IntCode[op3] = 1
			} else {
				IntCode[op3] = 0
			}

			// Increment
			pc += 4
		} else {
			panic("Unknown opcode")
		}

	}
}

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
	code = runComputer(code)
	fmt.Println("%#v", code)
}
