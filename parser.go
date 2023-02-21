package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// inputSplit find operation symbol and split string using it
func inputSplit(input string) (x, y, operation string) {
	for _, v := range []rune(input) {
		if v == 42 || v == 43 || v == 45 || v == 47 {
			operation = string(v)
			break
		}
	}

	values := strings.Split(input, operation)
	x = values[0]
	y = values[1]

	return
}

// inputCheck check entered data for major errors like:
//   - to many operators where entered
//   - as roman so arabic values entered
//   - no roman no arabic values entered
//     wrong operator
func inputCheck(input string) (bool, error) {
	var (
		isArabic      = false
		isRoman       = false
		noOperator    = true
		operatorCount = 0
	)

	for _, v := range []rune(input) {
		if unicode.IsDigit(v) {
			isArabic = true
		}
		if unicode.IsLetter(v) {
			isRoman = true
		}
		if v == 42 || v == 43 || v == 45 || v == 47 {
			operatorCount++
			if operatorCount > 1 {
				return false, fmt.Errorf("to many operators")
			}
			noOperator = false
		}
	}

	if isArabic && isRoman {
		return false, fmt.Errorf("different number systems are used")
	}
	if noOperator || (!isArabic && !isRoman) {
		return false, fmt.Errorf("that is not mathematical operation")
	}

	return isArabic, nil
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
