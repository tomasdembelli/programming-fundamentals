# Programming Fundamentals

This repository explains the programming fundamentals with some practical examples.

In computing, [source code](https://en.wikipedia.org/wiki/Source_code) is any collection of code written using a human-readable programming language, usually as plain text. The source code of a program is specially designed to facilitate the work of computer programmers, who specify the actions to be performed by a computer mostly by writing source code. The source code is often transformed by an assembler or compiler into binary machine code that can be executed by the computer. The machine code might then be stored for execution at a later time.

A [compiler](https://en.wikipedia.org/wiki/Compiler) is a computer program that translates computer code written in one programming language (the source language) into another language (the target language). The name "compiler" is primarily used for programs that translate source code from a high-level programming language to a lower level language (e.g. assembly language, object code, or machine code) to create an executable program.

In computer programming, [machine code](https://en.wikipedia.org/wiki/Machine_code) is any low-level programming language, consisting of machine language instructions, which are used to control a computer's central processing unit (CPU). Each instruction causes the CPU to perform a very specific task, such as a load, a store, a jump, or an arithmetic logic unit (ALU) operation on one or more units of data in the CPU's registers or memory.

Following is the general outline and the flowchart of a program that runs health check for a given list of URLs. 

## Outline

Outline of the health check program.
![Diagram](./media/Programming%20Fundamentals.svg)

## Flowchart

A flowchart* is a type of diagram that represents a workflow or process. A flowchart can also be defined as a diagrammatic representation of an algorithm, a step-by-step approach to solving a task.

Flowchart of the health check algorithm.
![Flowchart](./media/HealthCheck.svg)

*Flowchart [wiki](https://en.wikipedia.org/wiki/Flowchart).

[Iodraw](https://www.iodraw.com/diagram/) online flowchart application.


## Pseudo Code

Pseudo code is a term which is often used in programming and algorithm based fields. It is a methodology that allows the programmer to represent the implementation of an algorithm. [Source](https://www.geeksforgeeks.org/how-to-write-a-pseudo-code/)

Pseudo code for the health check program.
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

A Go program running health check.  Placeholder functions are used for the sake of simplicity.

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