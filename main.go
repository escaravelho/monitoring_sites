package main

import (
	"fmt"
	"os"
)

func main() {

	showIntro()

	showMenu()

	command := captureCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Show logs...")
	case 0:
		fmt.Println("Exit program...")
		os.Exit(0)
	default:
		fmt.Println("Error! Option choosed is unavailable.")
		os.Exit(-1)
	}
}

func showIntro() {
	version := 1.0
	fmt.Println("Monitoring sites - version: ", version)
	fmt.Println()
}

func showMenu() {
	fmt.Println("Initiating monitoring sites program...")
	fmt.Println("1- Initiating monitor")
	fmt.Println("2- Show logs")
	fmt.Println()
	fmt.Println("0- Exit program")
	fmt.Println()
}

func captureCommand() int {
	command := 0
	fmt.Println("What's your command to the program: ")
	fmt.Scan(&command)
	fmt.Println("The command that you choosed is", command)
	return command
}
