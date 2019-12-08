package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
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

func runComputer(IntCode []int, readcbk func() int, writecbk func(int)) []int {
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
			IntCode[IntCode[pc+1]] = readcbk()

			// Increment
			pc += 2
		} else if opCode[0] == OpOut {
			// Read
			op1 := getOperand(IntCode, pc+1, opCode[1])

			// Execute
			writecbk(op1)

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

var ampIndex int
var ampReads int
var ampPhase [5]int
var ampInput [5]int
var ampOutput [5]int

func read() int {
	ampReads++
	if ampReads == 1 {
		return ampPhase[ampIndex]
	} 
	return ampInput[ampIndex]
}

func write(value int) {
	// Save it
	ampOutput[ampIndex] = value

	// Print it out
	//fmt.Println(value)
}

func main() {
	// Get the input code
	orgInput := getInput()

	// Search for the best configuration
	maxThrust := 0
	bestPhase := make([]int, 5)
	for a := 0; a < 5; a++ {
		ampPhase[0] = a
		for b := 0; b < 5; b++ {
			ampPhase[1] = b
			for c := 0; c < 5; c++ {
				ampPhase[2] = c
				for d := 0; d < 5; d++ {
					ampPhase[3] = d
					for e := 0; e < 5; e++ {
						ampPhase = {a, b, c, d, e}
						
						// Reset the globals
						for i := 0; i < 5; i++ {
							ampInput[i] = 0
							ampOutput[i] = 0
						}

						// Setup the 5 amplifiers
						code := make([][]int, 5)
						for i := 0; i < 5; i++ {
							for j := 0; j < len(orgInput); j++ {
								code[i] = append(code[i], orgInput[j])
							}
						}

						// Run each of the amplifiers
						for i := 0; i < 5; i++ {
							// Reset the indices for each 
							ampIndex = i
							ampReads = 0

							// Run the computer
							code[i] = runComputer(code[i], read, write)

							// Copy the output
							if i < 4 {
								ampInput[i+1] = ampOutput[i]
							}
						}

						// Check for performance
						if ampOutput[4] > maxThrust {
							maxThrust = ampOutput[4]
							for i := 0; i < 5; i++ {
								bestPhase[i] = ampPhase[i]
							}
						}
					}
				}
			}
		}
	}
	
	fmt.Println("%d: %#v", maxThrust, bestPhase)
}
