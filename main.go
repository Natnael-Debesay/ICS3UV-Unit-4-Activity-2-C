// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-12-18
// Fileoverview: This program uses a While loop, with strings.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	// variables
	var someString string
	var someStringUpper string

	var reader = bufio.NewReader(os.Stdin)

	// input some string from the user
	fmt.Print("Please enter a string (type 'exit' to quit): ")
	someString, _ = reader.ReadString('\n')
	someString = strings.TrimSpace(someString)
	someStringUpper = strings.ToUpper(someString)

	// keep looking until the user enters 'exit' to quit
	for someStringUpper != "EXIT" {
		// output the string
		fmt.Println("You entered: " + someString)

		// ask for the next string
		fmt.Print("Please enter a string (type 'exit' to quit): ")
		someString, _ = reader.ReadString('\n')
		someString = strings.TrimSpace(someString)
		someStringUpper = strings.ToUpper(someString)
	}

	fmt.Println("\nDone.")
}
