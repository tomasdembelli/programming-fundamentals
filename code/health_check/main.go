package main

import (
	"bufio"
	"fmt"
	"os"
)

func extractURLs(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		return []string{}
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}

	var URLs []string = []string{
		"http://serviceA",
		"http://serviceB",
	}
	return URLs
}

func isHealthy(URL string) bool {
	var HTTPResponse int
	HTTPResponse = 404

	if HTTPResponse == 200 {
		return true
	} else {
		return false
	}
}

func main() {
	var listOfURL []string
	listOfURL = extractURLs("input_file_path")

	if len(listOfURL) == 0 {
		fmt.Println("URL list can not be empty")
		return
	}

	var healthyServices []string
	var unhealthyServices []string

	for n := 0; n < len(listOfURL); n++ {
		if isHealthy(listOfURL[n]) {
			healthyServices = append(healthyServices, listOfURL[n])
		} else {
			unhealthyServices = append(unhealthyServices, listOfURL[n])
		}
	}

	fmt.Println("Healthy services", healthyServices)
	fmt.Println("Unhealthy services", unhealthyServices)
}