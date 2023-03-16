package parser

import (
	"fmt"
	"strings"
)

func ParseString(input string, operators *[]rune) (*[]string, *[]string, error) {
	var (
		operatorSlice, operandSlice []string
		isOperator                  = false
		stringCopy                  = input
		operand                     string
	)

	for _, value := range []rune(input) {
		if !strings.ContainsRune(string(*operators), value) {
			isOperator = false
			continue
		} else if isOperator {
			return nil, nil, fmt.Errorf("to many operators")
		}
		operand, stringCopy, _ = strings.Cut(stringCopy, string(value))
		operatorSlice = append(operatorSlice, string(value))
		operandSlice = append(operandSlice, operand)
		isOperator = true
	}

	operandSlice = append(operandSlice, stringCopy)
	return &operandSlice, &operatorSlice, nil
}
