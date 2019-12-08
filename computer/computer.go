package computer

// Msg defines the message structure used by this computer
type Msg struct {
	Sender string
	Data   int
}

// Computer defines a new computer
type Computer struct {
	Input      chan Msg
	LogChannel *chan string
	Name 	   string
	Output     chan Msg
	

	code []int
	pc int
	outputs    []int
}

const (
	maxOpSize = 4
	maxNumOperands = (maxOpSize - 1)
)

const (
	paramModePosition  = 0
	paramModeImmediate = 1
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
	opHalt      = 99
)

// NewComputerWithName creates a new computer with a name
func NewComputerWithName(name string, code []int) Computer {
	memCopy := make([]int, len(code))

	copy(memCopy, code)

	return Computer{
		Input: make(chan Msg),
		Output: make(chan Msg),
		Name: name,

		code:   memCopy,
		outputs: make([]int, 0),
		pc: 0,
	}
}

// Run runs the computer
func (c *Computer) Run() {
	c.pc = 0

	// Go into an infinite loop
	for {
		// Get the Op Code
		opCode, modes := c.decodeOp()

		// Check for halt
		if opCode == opHalt {
			break
		}

		if opCode == opAdd {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])
			op3 := c.code[c.pc+3]

			// Execute
			c.code[op3] = op1 + op2

			// Increment the program counter
			c.pc += 4
		} else if opCode == opMul {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])
			op3 := c.code[c.pc+3]

			// Execute
			c.code[op3] = op1 * op2

			// Increment
			c.pc += 4
		} else if opCode == opIn {
			// Read
			msg := <-c.Input

			// Execute
			c.code[c.code[c.pc+1]] = msg.Data

			// Increment
			c.pc += 2
		} else if opCode == opOut {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])

			// Execute
			c.trySend(op1)
			c.outputs = append(c.outputs, op1)

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
			op3 := c.code[c.pc+3]

			// Execute
			if op1 < op2 {
				c.code[op3] = 1
			} else {
				c.code[op3] = 0
			}

			// Increment
			c.pc += 4
		} else if opCode == opEquals {
			// Read
			op1 := c.getOperand(c.pc+1, modes[0])
			op2 := c.getOperand(c.pc+2, modes[1])
			op3 := c.code[c.pc+3]

			// Execute
			if op1 == op2 {
				c.code[op3] = 1
			} else {
				c.code[op3] = 0
			}

			// Increment
			c.pc += 4
		} else {
			panic("Unknown opcode")
		}

	}
}

// GetLastOutput gets the last output by this computer
func (c *Computer) GetLastOutput() int {
	return c.outputs[len(c.outputs)-1]
}

func (c *Computer) getOperand(pos int, mode int) int {
	if mode == paramModePosition {
		return c.code[c.code[pos]]
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