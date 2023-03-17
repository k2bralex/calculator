package app

import (
	. "calc/internal/inputhandler"
	. "calc/internal/model"
	. "calc/internal/pkg"
	"fmt"
)

// Run calculator
func Run() {
	for {
		// Get input and if "exit" terminates execution
		fmt.Println("Input expression:")
		input := InputReader()
		if input == "exit" {
			return
		}

		opd, ops, err := InputHandler(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		c := Calculator{Operands: opd, Operators: ops}
		fmt.Println("Result: " + c.Start())
	}
}
