package main

import "fmt"

// OP Codes
const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
)

var inputValues = []uint32{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 19, 6, 23, 2, 23, 13, 27, 1, 27, 5, 31, 2, 31, 10, 35, 1, 9, 35, 39, 1, 39, 9, 43, 2, 9, 43, 47, 1, 5, 47, 51, 2, 13, 51, 55, 1, 55, 9, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 10, 67, 71, 1, 71, 10, 75, 2, 75, 13, 79, 2, 79, 13, 83, 1, 5, 83, 87, 1, 87, 6, 91, 2, 91, 13, 95, 1, 5, 95, 99, 1, 99, 2, 103, 1, 103, 6, 0, 99, 2, 14, 0, 0}

func runComputer(IntCode []uint32) bool {
	var pc uint32 = 0
	for {
		var opCode uint32 = IntCode[pc]

		if opCode == OpHalt {
			break
		}

		var op1 uint32 = IntCode[pc+1]
		var op2 uint32 = IntCode[pc+2]
		var op3 uint32 = IntCode[pc+3]

		if op1 > uint32(len(IntCode)) ||
			op2 > uint32(len(IntCode)) ||
			op3 > uint32(len(IntCode)) {
			return false
		}

		if opCode == OpAdd {
			IntCode[op3] = IntCode[op1] + IntCode[op2]
		} else {
			IntCode[op3] = IntCode[op1] * IntCode[op2]
		}

		/* Increment the PC */
		pc = pc + 4
	}

	return true
}

func main() {
	tempValues := make([]uint32, len(inputValues))

	/* Go into a brute force search */
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			/* Create the temporary values */
			for i := 0; i < len(inputValues); i++ {
				tempValues[i] = inputValues[i]
			}

			/* Make the changes */
			tempValues[1] = uint32(noun)
			tempValues[2] = uint32(verb)

			/* Run the computer */
			success := runComputer(tempValues)

			/* Evaluate */
			if success {
				if tempValues[0] == 19690720 {
					fmt.Println(noun, " ", verb)
					break
				}
			}
		}
	}
}
