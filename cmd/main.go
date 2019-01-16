package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	parsing "github.com/andskur/number-parsing"
)

func main() {
	// Infinity loop for accepting user integer string from stdin
	for {
		// Accepting string from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Input an Integer:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSuffix(input, "\n")

		// Check if inout is empty
		if len(input) < 1 {
			fmt.Println(errors.New("empty string"))
			continue
		}

		// Convert string to int
		number, err := strconv.Atoi(input)
		if err != nil {
			// handle the error
			fmt.Println(fmt.Sprintf("you can enter only numbers between %d and %d", parsing.MinInt, parsing.MaxInt))
			continue
		}

		// convert given int to human-readable english
		output := parsing.ConvertNumbers(number)

		// print getting value to stdout
		fmt.Println(output + "\n")
	}
}