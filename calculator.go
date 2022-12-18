package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter a number: ")
	var number1 int
	fmt.Scan(&number1)

	fmt.Print("Enter another number: ")
	var number2 int
	fmt.Scan(&number2)

	fmt.Print("Enter add,subract,multiply or divide: ")
	var operation string
	fmt.Scan(&operation)

	sum := number1 + number2
	diff := number1 - number2
	prod := number1 * number2
	div := number1 / number2

	switch operation {
	case "add":
		fmt.Print("Sum: ", sum)
	case "subract":
		fmt.Print("Difference: ", diff)
	case "multiply":
		fmt.Print("Product: ", prod)
	case "divide":
		fmt.Print("Quotient: ", div)

	}
}
