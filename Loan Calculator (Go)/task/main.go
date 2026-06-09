package main

import (
	"flag"
	"fmt"
	"math"
)

func formatPeriods(months int) string {
	years := months / 12
	rem := months % 12

	yearStr := func(y int) string {
		if y == 1 {
			return "1 year"
		}
		return fmt.Sprintf("%d years", y)
	}
	monthStr := func(m int) string {
		if m == 1 {
			return "1 month"
		}
		return fmt.Sprintf("%d months", m)
	}

	switch {
	case years == 0:
		return monthStr(rem)
	case rem == 0:
		return yearStr(years)
	default:
		return yearStr(years) + " and " + monthStr(rem)
	}
}

func main() {
	principal := flag.Float64("principal", 0, "")
	periods := flag.Int("periods", 0, "")
	payment := flag.Float64("payment", 0, "")
	interest := flag.Float64("interest", 0, "")
	flag.Parse()

	i := *interest / (12 * 100)

	switch {
	case *payment == 0:
		// calculate annuity payment
		p := *principal
		n := float64(*periods)
		pow := math.Pow(1+i, n)
		a := math.Ceil(p * i * pow / (pow - 1))
		fmt.Printf("Your monthly payment = %.0f!\n", a)

	case *periods == 0:
		// calculate number of months
		a := *payment
		p := *principal
		n := math.Ceil(math.Log(a/(a-i*p)) / math.Log(1+i))
		fmt.Printf("It will take %s to repay this loan!\n", formatPeriods(int(n)))

	case *principal == 0:
		// calculate loan principal
		a := *payment
		n := float64(*periods)
		pow := math.Pow(1+i, n)
		p := math.Floor(a / (i * pow / (pow - 1)))
		fmt.Printf("Your loan principal = %.0f!\n", p)
	}
}
