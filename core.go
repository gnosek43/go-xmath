package xmath

import (
	"errors"
	"fmt"
)

func PrintHello() {
	fmt.Println("Hello, Modules! This is xmath speaking!")
}

func Add(a, b int) int {
	return a + b
}

func Div(a, b int) (error, float64) {
	if b == 0 {
		return errors.New("can't divide by 0"), 0
	}
	return nil, float64(a) / float64(b)
}

func Pow(a, b int) int {
	base := a
	for i := 1; i < b; i++ {
		a = a * base
	}
	return a
}
