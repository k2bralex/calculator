package main

import "fmt"

// Run calculator
func Run() {
	for {
		// Get input and if "exit" terminates execution
		fmt.Println("Input expression:")
		input := inputReader()
		if input == "exit" {
			return
		}

		// Check user input and print if any error and go to next iteration
		isArabic, err := inputCheck(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Get operands and operator
		xStr, yStr, op := inputSplit(input)
		var x, y int

		// Parse spitted strings to int values due to roman/arabic input
		if isArabic {
			x, err = valueParse(xStr)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			y, err = valueParse(yStr)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		} else {
			x, err = romanToArabic(xStr)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			y, err = romanToArabic(yStr)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}

		// Calculate result
		var result int
		switch op {
		case "+":
			result = calc(x, y, add)
		case "-":
			result = calc(x, y, sub)
		case "*":
			result = calc(x, y, mult)
		case "/":
			result = calc(x, y, div)
		}

		// Write result to console
		if isArabic {
			fmt.Printf("%d\n", result)
		} else {
			r, err := arabicToRoman(result)
			romanPrinter(r, err)
		}

	}
}

func romanPrinter(s string, err error) {
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%s\n", s)
}
