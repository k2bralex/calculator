package app

import (
	"bufio"
	. "calc/internal/inputhandler"
	"calc/internal/model"
	"fmt"
	"os"
)

// Run calculator
func Run() {
	for {
		// Get input and if "exit" terminates execution
		fmt.Println("Input expression:")
		input := inputReader()
		if input == "exit" {
			return
		}

		opd, ops, err := InputHandler(input)
		if err != nil {
			fmt.Println(err)
		}

		c := model.Calculator{Operands: opd, Operators: ops}
		fmt.Println("Result: " + c.Start())
	}
}

func romanPrinter(s string, err error) {
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%s\n", s)
}

// inputReader use bufio Package read from console and save into string
func inputReader() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return ""
	}

	return scanner.Text()
}
