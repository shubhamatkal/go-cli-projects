package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags for operation and operands
	operation := flag.String("op", "", "Operation to perform: add, sub, mul, div")
	num1 := flag.Float64("num1", 0, "First number")
	num2 := flag.Float64("num2", 0, "Second number")

	// Parse the flags
	flag.Parse()

	// Ensure an operation is provided
	if *operation == "" {
		fmt.Println("Error: Please specify an operation using the -op flag.")
		fmt.Println("Example: go run main.go -op=add -num1=3 -num2=4")
		os.Exit(1)
	}

	// Perform the operation
	switch *operation {
	case "add":
		fmt.Printf("Result: %.2f\n", *num1+*num2)
	case "sub":
		fmt.Printf("Result: %.2f\n", *num1-*num2)
	case "mul":
		fmt.Printf("Result: %.2f\n", *num1**num2)
	case "div":
		if *num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", *num1 / *num2)
	default:
		fmt.Println("Error: Unsupported operation. Use add, sub, mul, or div.")
		os.Exit(1)
	}
}
