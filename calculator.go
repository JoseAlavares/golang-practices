package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
)

func contains(value string, array []string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}

	return false
}

func loopOperation(operationType string, values []int) int {
	var result int = 0
	for _, number := range values {
		if operationType == "+" {
			result += number
		} else if operationType == "-" {
			result -= number
		}
	}

	return result
}

func scanInputs(scanner bufio.Scanner) ([]int, error) {
	var counter int = 2
	var number string
	var valuesToIterate []int

	for i:=0; i<counter; i++ {
		if i == 0 {
			number = "first"
		} else {
			number = "second"
		}

		fmt.Println("Enter the",  number,  " number")

		scanner.Scan()
		inputStr := scanner.Text()
		input, error := strconv.Atoi(inputStr)

		if error != nil {
			fmt.Println("Error in the inputs")
			return nil, errors.New("Error in the inputs")
		}

		valuesToIterate = append(valuesToIterate, input)
	}

	return valuesToIterate, nil
}

func main() {
	var sentinel bool = true
	operationsAllowed := []string{ "+", "-" }
	scanner := bufio.NewScanner(os.Stdin)

	for sentinel {
		fmt.Println("Please select what type of operation you need\n+ -> sum\n- -> subtraction")

		scanner.Scan()
		selectedOperation := scanner.Text()

		if !contains(selectedOperation, operationsAllowed) {
			fmt.Println("Operación no permitida")
			continue
		} else if selectedOperation == "close" {
			break
		}

		inputsToIterate, error := scanInputs(*scanner)

		if error != nil {
			fmt.Println("Error in the function scanInputs")
			return
		}

		result := loopOperation(selectedOperation, inputsToIterate)
		fmt.Println("Result", result)
	}

}