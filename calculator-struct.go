package main

import (
	"bufio"
	"os"
)

scanner := bufio.NewScanner(os.Stdin)
allowedInputs := 2
allowedOperations := 4
divisionAllowedOperations := 1

type Calculator struct {
	inputs []int
	operator string
	result float64
}

func (c Calculator) validateOperator(operator string) bool {
	allowedOperators := [4]string {"+", "-", "*", "/"}

	for operator_ := range allowedOperators {
		if operator_ == operator {
			return true
		}
	}

	return false
}

func captureInputs(operator string, input int) []int {
	var inputs []int

	for i:=0; i<allowedInputs; i++ {
		
	}
}

func (c *Calculator) operations() {

}

func iterateInputs() []int {
	 
}

func main() {

}