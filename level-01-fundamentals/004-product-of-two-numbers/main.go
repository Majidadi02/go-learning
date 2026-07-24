package main

import "fmt"

func main() {
	fmt.Print("Enter the first number: ")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Print("Enter the second number: ")
	var secondNumber int
	fmt.Scan(&secondNumber)
	product := firstNumber * secondNumber
	fmt.Println("========================")
	fmt.Println("    PRODUCT CALCULATOR    ")
	fmt.Println("========================")
	fmt.Println("The first number is:", firstNumber)
	fmt.Println("The second number is:", secondNumber)
	fmt.Println("Product:", product)
	fmt.Println("    Bonus Challenge   ")
	fmt.Println(firstNumber, "x", secondNumber, "=", product)
}
