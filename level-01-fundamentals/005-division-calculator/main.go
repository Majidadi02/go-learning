package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the first number: ")
	var dividend int
	fmt.Scan(&dividend)
	fmt.Print("Enter the second number: ")
	var divisor int
	fmt.Scan(&divisor)
	quotient := dividend / divisor
	fmt.Println("========================")
	fmt.Println("    DIVISION CALCULATOR    ")
	fmt.Println("========================")
	fmt.Println("Dividend:", dividend)
	fmt.Println("Divisor:", divisor)
	fmt.Println("Quotient:", quotient)
	fmt.Println("    Bonus Challenge   ")
	remainder := dividend % divisor
	fmt.Println("Remainder: ", remainder)
}
