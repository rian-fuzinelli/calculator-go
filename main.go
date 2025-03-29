package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Calc with Go")
	fmt.Println("Choose the operator: +, -, *, /")
	fmt.Scanln(&operator)

	fmt.Print("Type the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Type the second number: ")
	fmt.Scanln(&num2)

	var result float64
	var error bool

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			error = true
			fmt.Println("error: division by zero!")
		} else {
			result = num1 / num2
		}
	default:
		error = true
		fmt.Println("Invalid operation!")
	}

	if !error {
		fmt.Printf("result: %.0f\n", result)
	}
}
