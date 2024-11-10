package main

import (
	"fmt"
)

type Calculator interface {
	Calculate(float64, float64) float64
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}

func (a Addition) Calculate(x, y float64) float64 {
	return x + y
}

func (s Subtraction) Calculate(x, y float64) float64 {
	return x - y
}

func (m Multiplication) Calculate(x, y float64) float64 {
	return x * y
}

func (d Division) Calculate(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error! Division by zero.")
		return 0
	}
	return x / y
}

func main() {
	fmt.Println("Select operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	var choice string
	fmt.Print("Enter choice (1/2/3/4): ")
	fmt.Scanln(&choice)

	var num1, num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	var calculator Calculator
	switch choice {
	case "1":
		calculator = Addition{}
	case "2":
		calculator = Subtraction{}
	case "3":
		calculator = Multiplication{}
	case "4":
		calculator = Division{}
	default:
		fmt.Println("Invalid input")
		return
	}

	result := calculator.Calculate(num1, num2)
	fmt.Println("Result:", result)

	fmt.Print("Do you want to continue? (y/n): ")
		var continueChoice string
		fmt.Scanln(&continueChoice)

		if continueChoice != "y" {
			keepCalculating = false
		}

		// Обработка сигнала прерывания
		select {
		case <-c:
			fmt.Println("Exiting...")
			return
		default:
		}
	}
}
