package model

import (
	"calc/internal/converter"
	"strconv"
)

type Calculator struct {
	Operands  interface{}
	Operators *[]string
}

func (c *Calculator) Start() string {
	var (
		result int64
	)

	arabic, ok := (c.Operands).(*[]int64)
	if !ok {
		arabic = &[]int64{}
		for _, v := range *(c.Operands).(*[]string) {
			val, err := converter.Decode(v)
			if err != nil {
				return err.Error()
			}
			*arabic = append(*arabic, val)
		}
	}

	result = Calculate(arabic, c.Operators)

	if !ok {
		res, err := converter.Encode(result)
		if err != nil {
			return err.Error()
		}
		return res
	}

	return strconv.FormatInt(result, 10)
}

func Calculate(nums *[]int64, op *[]string) int64 {
	var (
		opFunc func(int64, int64) int64
		result = (*nums)[0]
	)

	for i := 1; i < len(*nums); i++ {
		switch (*op)[i-1] {
		case "+":
			opFunc = add
		case "-":
			opFunc = sub
		case "*":
			opFunc = mult
		case "/":
			opFunc = div
		}

		result = handler(result, (*nums)[i], opFunc)
	}

	return result
}

func handler(x, y int64, op func(int64, int64) int64) int64 {
	return op(x, y)
}

func add(x, y int64) int64  { return x + y }
func sub(x, y int64) int64  { return x - y }
func mult(x, y int64) int64 { return x * y }
func div(x, y int64) int64 {
	if y == 0 {
		return 0
	}
	return x / y
}
