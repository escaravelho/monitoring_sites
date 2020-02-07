package main

import "fmt"

func main() {

	version := 1.0
	fmt.Println("Monitoring sites - version: ", version)
	fmt.Println()
	fmt.Println("Initiating monitoring sites program...")
	fmt.Println("1- Initiating monitor")
	fmt.Println("2- Show logs")
	fmt.Println()
	fmt.Println("0- Exit program")
	fmt.Println()

	command := 0
	fmt.Println("What's your command to the program: ")
	fmt.Scan(&command)
	fmt.Println("The command that you choosed is", command)
}
