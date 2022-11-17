package xmath

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is xmath speaking!")
}

func Add(a, b int) int {
	return a + b
}

func Div(a, b int) float64 {
	return float64(a / b)
}

func Pow(a, b int) int {
	base := a
	for i := 1; i < b; i++ {
		a = a * base
	}
	return a
}
