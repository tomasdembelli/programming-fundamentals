# Programming Fundamentals

This repository explains the programming fundamentals with some practical examples.

Following is the general outline and the flowchart of a program that runs health check for a given list of URLs. 

## Outline
![Diagram](./media/Programming%20Fundamentals.svg)

## Flowchart
A flowchart* is a type of diagram that represents a workflow or process. A flowchart can also be defined as a diagrammatic representation of an algorithm, a step-by-step approach to solving a task.
![Flowchart](./media/HealthCheck.svg)

*Flowchart [wiki](https://en.wikipedia.org/wiki/Flowchart).

[Iodraw](https://www.iodraw.com/diagram/) online flowchart application.

## Pseudo Code

Pseudo code is a term which is often used in programming and algorithm based fields. It is a methodology that allows the programmer to represent the implementation of an algorithm. [Source](https://www.geeksforgeeks.org/how-to-write-a-pseudo-code/)

```
function extractURLs(filePath) {
    Open file given by the filePath.
    Extract URLs.(Regex, line by line, ..)
    Deduplicate the list.
    ...
    return a slice of string which each element is a URL to a service
}

function isHealthy(url) {
    Send an HTTP request to the given URL.
    If the response status code is 200, then return true.
    Otherwise, return false.
}

function main() {
    Acquire the input file.
    Extract the URLs via exctractURLs function.
    Validate the health of each service via isHealthy function.
    Store the responses as healthy and unhealthy services.
    Print out the results.
}
```



## Code

A Go program with placeholder functions for the sake of simplicity.

```Go
package main

import "fmt"

func extractURLs() []string {
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
	listOfURL = extractURLs()

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
```