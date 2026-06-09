package main

import (
	"fmt"
	"math"
)

func main() {
	var principal float64
	fmt.Println("Enter the loan principal:")
	fmt.Scan(&principal)

	fmt.Println("What do you want to calculate?")
	fmt.Println("type \"m\" for number of monthly payments,")
	fmt.Print("type \"p\" for the monthly payment:\n")

	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "m":
		var payment float64
		fmt.Println("Enter the monthly payment:")
		fmt.Scan(&payment)
		months := int(math.Ceil(principal / payment))
		if months == 1 {
			fmt.Println("It will take 1 month to repay the loan")
		} else {
			fmt.Printf("It will take %d months to repay the loan\n", months)
		}
	case "p":
		var months float64
		fmt.Println("Enter the number of months:")
		fmt.Scan(&months)
		payment := int(math.Ceil(principal / months))
		lastPayment := int(principal) - (int(months)-1)*payment
		if lastPayment == payment {
			fmt.Printf("Your monthly payment = %d\n", payment)
		} else {
			fmt.Printf("Your monthly payment = %d and the last payment = %d.\n", payment, lastPayment)
		}
	}
}
