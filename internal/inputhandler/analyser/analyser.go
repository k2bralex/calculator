package pkg

import (
	. "calc/internal/pkg"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func OperandCheck(opd *[]string) (*[]string, *[]int64, error) {
	var (
		isDigit  = false
		isLetter = false
	)
	var nums []int64
	for _, str := range *opd {
		value, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			isLetter = true
		} else {
			nums = append(nums, value)
			isDigit = true
		}
		if isDigit && isLetter {
			return nil, nil, fmt.Errorf("different value systems are used")
		}
	}
	if isLetter {
		return opd, nil, nil
	}
	return nil, &nums, nil
}

func OperatorsCheck(input string, operators *[]rune) error {
	if strings.ContainsAny(string(*operators), input[0:1]) ||
		strings.ContainsAny(string(*operators), input[len(input)-1:]) {
		return fmt.Errorf("not mathematical expression")
	}
	for _, r := range []rune(input) {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && !Contains(r, *operators) {
			return fmt.Errorf("unknown operator")
		}
	}
	return nil
}
