package main

import (
	"goBfInterpreter/bfinterpreter"
	"os"
)

func main() {
	code := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	machine := bfinterpreter.NewBfMachine(code, os.Stdin, os.Stdout)
	machine.Exec()
}