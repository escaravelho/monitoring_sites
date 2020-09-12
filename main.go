package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {

	showIntro()

	for {
		showMenu()

		command := captureCommand()

		switch command {
		case 1:
			startMonitor()
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

func startMonitor() {
	fmt.Println("Monitoring...")

	sites := [] string{"https://www.alura.com.br", "https://www.caelum.com.br", "https://www.campuscode.com.br"}

	for _, site := range sites {
		
		response, _ := http.Get(site)
		
		if response.StatusCode == 200 {
			fmt.Println("The website ", site, "is works fine!")
		} else {
			fmt.Println("The website ", site, "is not work!")
		}
	}
}