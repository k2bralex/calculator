package main

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int  { return x + y }
func sub(x, y int) int  { return x - y }
func mult(x, y int) int { return x * y }
func div(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}
