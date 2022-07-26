package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiple() int
	Division() float64
	Substract() int
}

func (n *Numbers) Sum() int {
	return n.num1 + n.num2
}
func (n *Numbers) Multiplem() int {
	return n.num1 * n.num2
}
func (n *Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}
func (n *Numbers) Substract() int {
	return n.num1 - n.num2
}

func main() {
	numbers := Numbers{9, 8}
	fmt.Println(numbers.Multiplem())
}
