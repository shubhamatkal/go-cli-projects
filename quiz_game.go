package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Quiz game!")
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, Welcome to the game\n", name)
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	fmt.Println(age >= 10)
	if age >= 10 {
		fmt.Println("yay you can play")
	} else {
		fmt.Println("you cannot play!")
		return
	}
	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct")

	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("correct ")
	} else {
		fmt.Println("Incorrect")
	}

}
