package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Digite o calculo que você quer fazer. Ex.:(2*2)")
	var input string
	fmt.Scan(&input)
	operador := strings.Split(input, "")
	resp := calculadora(operador)
	fmt.Println(resp)
}

func calculadora(op []string) int {

	num1, _ := strconv.Atoi(op[0])
	num2, _ := strconv.Atoi(op[2])
	var result int

	switch op[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}
