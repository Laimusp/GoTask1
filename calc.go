package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите операцию (например, 3 + 5): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		fmt.Println("Error: Invalid input format")
		os.Exit(1)
	}

	operation := parts[1]
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[2])

	var result int

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operation")
		os.Exit(1)
	}

	fmt.Printf("%d %s %d = %d\n", num1, operation, num2, result)
}
