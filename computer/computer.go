package computer

import (
	"strconv"
)

// Msg defines the message structure used by this computer
type Msg struct {
	Sender string
	Data   int
}

// Computer defines a new computer
type Computer struct {
	Input      chan Msg
	Name 	   string
	Output     chan Msg
	IsHalted   bool
	IsWaitingForInput bool
	IsReadyForOutput bool

	code       []int
	pc         int
	outputs    []int
	relBase	   int
}

const (
	maxOpSize = 4
	maxNumOperands = (maxOpSize - 1)
)

const (
	paramModePosition  = 0
	paramModeImmediate = 1
	paramModeRelative  = 2
)

const (
	opAdd       = 1
	opMul       = 2
	opIn        = 3
	opOut       = 4
	opJumpTrue  = 5
	opJumpFalse = 6
	opLessThan  = 7
	opEquals    = 8
	opAdjRel    = 9
	opHalt      = 99
)

// NewComputerWithName creates a new computer with a name
func NewComputerWithName(name string, code []int) Computer {
	memCopy := make([]int, 5000)

	copy(memCopy, code)

	log("Created: " + name)

	return Computer{
		Input: make(chan Msg),
		Output: make(chan Msg),
		Name: name,
		IsHalted: false,
		IsWaitingForInput: false,
		IsReadyForOutput: false,

		code:   memCopy,
		outputs: make([]int, 0),
		pc: 0,
		relBase: 0,
	}
}

func log(msg string) {
	// fmt.Println(msg)
}

// Run runs the computer
func (c *Computer) Run() {
	c.pc = 0
	c.IsHalted = false

	// Go into an infinite loop
	for {
		log(c.Name + " at PC: " + strconv.Itoa(c.pc))

		// Get the Op Code
		opCode, modes := c.decodeOp()

		// Check for halt
		if opCode == opHalt {
			c.IsHalted = true
			break
		}

		if opCode == opAdd {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			c.writeOperand(c.pc+3, modes[2], op1 + op2)

			// Increment the program counter
			c.pc += 4
		} else if opCode == opMul {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			c.writeOperand(c.pc+3, modes[2], op1 * op2)

			// Increment
			c.pc += 4
		} else if opCode == opIn {
			log(c.Name + " waiting for input")
			c.IsWaitingForInput = true

			// Read
			msg := <-c.Input
			log(c.Name + " got input: " + strconv.Itoa(msg.Data))

			// Execute
			c.IsWaitingForInput = false
			c.writeOperand(c.pc+1, modes[0], msg.Data)

			// Increment
			c.pc += 2
		} else if opCode == opOut {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			log(c.Name + " waiting to output")
			c.IsReadyForOutput = true

			// Execute
			c.trySend(op1)
			c.outputs = append(c.outputs, op1)
			log(c.Name + " Outputing " + strconv.Itoa(op1))
			c.IsReadyForOutput = false

			// Increment
			c.pc += 2
		} else if opCode == opJumpTrue {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			if op1 != 0 {
				c.pc = op2
			} else {
				// Increment
				c.pc += 3
			}
		} else if opCode == opJumpFalse {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			if op1 == 0 {
				c.pc = op2
			} else {
				// Increment
				c.pc += 3
			}
		} else if opCode == opLessThan {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			if op1 < op2 {
				c.writeOperand(c.pc+3, modes[2], 1)
			} else {
				c.writeOperand(c.pc+3, modes[2], 0)
			}

			// Increment
			c.pc += 4
		} else if opCode == opEquals {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])

			// Execute
			if op1 == op2 {
				c.writeOperand(c.pc+3, modes[2], 1)
			} else {
				c.writeOperand(c.pc+3, modes[2], 0)
			}

			// Increment
			c.pc += 4
		} else if opCode == opAdjRel {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])

			// Execute
			c.relBase += op1

			// Increment
			c.pc += 2
		} else {
			panic("Unknown opcode")
		}
	}
}

// GetLastOutput gets the last output by this computer
func (c *Computer) GetLastOutput() int {
	return c.outputs[len(c.outputs)-1]
}

func (c *Computer) writeOperand(pos int, mode int, value int) {
	if mode == paramModePosition {
		c.code[c.code[pos]] = value
	} else if mode == paramModeRelative {
		c.code[c.relBase + c.code[pos]] = value
	} else {
		panic("Unsupported write operand mode")
	}
}

func (c *Computer) getOperand(pos int, mode int) int {
	if mode == paramModePosition {
		return c.code[c.code[pos]]
	} else if mode == paramModeRelative {
		return c.code[c.relBase + c.code[pos]]
	}
	return c.code[pos]
}

func (c *Computer) decodeOp() (int, []int) {
	opCode := c.code[c.pc]

	// Get the actual op Code
	op := opCode % 100
	opCode /= 100

	// Get the modes
	modes := make([]int, maxNumOperands)
	for i := 0; i < maxNumOperands; i++ {
		if opCode <= 0 {
			break
		}
		val := opCode % 10
		modes[i] = val
		opCode /= 10
	}

	return op, modes
}

func (c *Computer) trySend(value int) {
	defer func() {
		recover() // nolint: errcheck
	}()

	msg := Msg{c.Name, value}
	if c.Output != nil {
		c.Output <- msg
	}
}