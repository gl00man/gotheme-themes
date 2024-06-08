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
	for {
		fmt.Print("Wprowadź pierwszą liczbę: ")
		num1Str, _ := reader.ReadString('\n')
		num1, err := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)
		if err != nil {
			fmt.Println("Niepoprawna liczba. Spróbuj ponownie.")
			continue
		}

		fmt.Print("Wprowadź operator (+, -, *, /): ")
		operator, _ := reader.ReadString('\n')
		operator = strings.TrimSpace(operator)

		fmt.Print("Wprowadź drugą liczbę: ")
		num2Str, _ := reader.ReadString('\n')
		num2, err := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)
		if err != nil {
			fmt.Println("Niepoprawna liczba. Spróbuj ponownie.")
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Błąd: Dzielenie przez zero.")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("Niepoprawny operator. Spróbuj ponownie.")
			continue
		}

		fmt.Printf("Wynik: %f\n", result)
	}
}
