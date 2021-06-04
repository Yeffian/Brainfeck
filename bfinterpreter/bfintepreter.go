package bfinterpreter

import (
	"io"
)

// Stores all the data our interpreter will need to use.
type BfMachine struct {
	code string
	instructionPtr int

	mem [3000]int
	dataPtr int

	buffer []byte

	input io.Reader
	output io.Writer
}

// Makes a new instance of the BfMachine struct.
func NewBfMachine(code string, input io.Reader, output io.Writer) *BfMachine {
	return &BfMachine {
		code: code,
		input: input,
		output: output,
		buffer: make([]byte, 1),
	}
}

// Executes our code.
func (m* BfMachine) Exec() {
	// Loop through all the instructions in the code
	for m.instructionPtr < len(m.code) {
		instruction := m.code[m.instructionPtr]  // Get the current instruction

		// Interpret the code
		switch instruction {
		// For + and -, we increment or decrement the memory at the data ptr,
		// but in case of < and >, we increment or decrement the data ptr itself.
		// [ needs to be accompanied with a ]. If the current cell contains a zero, set the instruction ptr to the index of the instruction after the matching [
		// and ] set the instruction ptr to the index of the instruction after the matching [
		case '+':
			m.mem[m.dataPtr]++
		case '-':
			m.mem[m.dataPtr]--
		case '>':
			m.dataPtr++
		case '<':
			m.dataPtr--
		case ',':
			m.ReadChar()
		case '.':
			m.PutChar()
		case '[':
			if m.mem[m.dataPtr] == 0 {
				depth := 1
				for depth != 0 {
					m.instructionPtr++
					switch m.code[m.instructionPtr] {
					case '[':
						depth++
					case ']':
						depth--
					}
				}
			}
		case ']':
			if m.mem[m.dataPtr] != 0 {
				depth := 1
				for depth != 0 {
					m.instructionPtr--
					switch m.code[m.instructionPtr] {
					case ']':
						depth++
					case '[':
						depth--
					}
				}
			}
		}

		// Increment the data ptr so we can move to the next instruction
		m.instructionPtr++
	}
}

// Reads a byte from the input and transfers that byte to the current memory cell.
func (m* BfMachine) ReadChar() {
	n, err := m.input.Read(m.buffer)

	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("Wrong num bytes read.")
	}

	m.mem[m.dataPtr] = int(m.buffer[0])
}

// Writes the current memory cell to os.Stdout
func (m* BfMachine) PutChar() {
	m.buffer[0] = byte(m.mem[m.dataPtr])

	n, err := m.output.Write(m.buffer)

	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("Wrong num bytes written.")
	}
}

