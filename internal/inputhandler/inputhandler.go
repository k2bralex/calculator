package inputhandler

import (
	. "calc/internal/inputhandler/analyser"
	. "calc/internal/inputhandler/parser"
)

var (
	OPERATORS = []rune{'+', '-', '*', '/'}
)

func InputHandler(input string) (interface{}, *[]string, error) {

	if err := OperatorsCheck(input, &OPERATORS); err != nil {
		return nil, nil, err
	}

	opd, opr, err := ParseString(input, &OPERATORS)
	if err != nil {
		return nil, nil, err
	}

	str, ints, err := OperandCheck(opd)
	if err != nil {
		return nil, nil, err
	}

	if ints == nil {
		return str, opr, nil
	}

	return ints, opr, nil
}
