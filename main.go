package main

import (
	"strings"
	"io"
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
)

const timesToMonitor = 2
const delayToNextMonitoring = 2

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

func readSitesFileToMonitor() [] string {
	
	var sites []string 
	file, error := os.Open("sites.properties")

	if error != nil {
		fmt.Println("Error to open sites file")
	}

	reader := bufio.NewReader(file)
	
	for {
		fileLine, error := reader.ReadString('\n')

		if error == io.EOF {
			break
		}
		
		fileLine = strings.TrimSpace(fileLine)
		sites = append(sites, fileLine)
	}
	file.Close()

	return sites
}

func startMonitor() {

	fmt.Println("Monitoring...")
	sites := readSitesFileToMonitor()

	for tentative := 1; tentative <= timesToMonitor; tentative++ {
		
		fmt.Println("Initializing the",tentative,"º tentative to monitor sites...\n")

		for _, site := range sites {
				
			response, error := http.Get(site)

			if error != nil {
				fmt.Println("Error to do a request to website", site)
			} else if response.StatusCode == 200 {
				fmt.Println("The website:",site,"is works fine!")
			} else {
				fmt.Println("The website:",site,"is not work!")
			}
			
			fmt.Println("Waiting",delayToNextMonitoring,"seconds to monitoring next site...")
			time.Sleep(delayToNextMonitoring * time.Second)
			fmt.Println("\n")
		}
	}  
}