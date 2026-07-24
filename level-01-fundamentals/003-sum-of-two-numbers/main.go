package main

import "fmt"

func main() {
	fmt.Print("Enter the first number: ")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Print("Enter the second number: ")
	var secondNumber int
	fmt.Scan(&secondNumber)
	sum := firstNumber + secondNumber
	fmt.Println("========================")
	fmt.Println("    SUM CALCULATOR    ")
	fmt.Println("========================")
	fmt.Println("The first number is:", firstNumber)
	fmt.Println("The second number is:", secondNumber)
	fmt.Println("Result:", sum)

}
