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

func incorrect() {
	fmt.Println("Incorrect parameters")
}

func main() {
	typeFlag := flag.String("type", "", "")
	principal := flag.Float64("principal", 0, "")
	periods := flag.Int("periods", 0, "")
	payment := flag.Float64("payment", 0, "")
	interest := flag.Float64("interest", 0, "")
	flag.Parse()

	// Count explicitly provided flags
	provided := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		provided[f.Name] = true
	})

	// Validation
	if *typeFlag != "annuity" && *typeFlag != "diff" {
		incorrect()
		return
	}
	if !provided["interest"] || *interest <= 0 {
		incorrect()
		return
	}
	if *typeFlag == "diff" && provided["payment"] {
		incorrect()
		return
	}
	if len(provided) < 4 {
		incorrect()
		return
	}
	if *principal < 0 || *payment < 0 || *interest < 0 || *periods < 0 {
		incorrect()
		return
	}

	i := *interest / (12 * 100)

	switch *typeFlag {
	case "diff":
		p := *principal
		n := *periods
		total := 0.0
		for m := 1; m <= n; m++ {
			dm := math.Ceil(p/float64(n) + i*(p-p*float64(m-1)/float64(n)))
			total += dm
			fmt.Printf("Month %d: payment is %.0f\n", m, dm)
		}
		fmt.Printf("\nOverpayment = %.0f\n", total-p)

	case "annuity":
		p := *principal
		n := float64(*periods)
		a := *payment

		switch {
		case !provided["payment"]:
			pow := math.Pow(1+i, n)
			a = math.Ceil(p * i * pow / (pow - 1))
			overpayment := n*a - p
			fmt.Printf("Your annuity payment = %.0f!\n", a)
			fmt.Printf("Overpayment = %.0f\n", overpayment)

		case !provided["periods"]:
			n = math.Ceil(math.Log(a/(a-i*p)) / math.Log(1+i))
			overpayment := n*a - p
			fmt.Printf("It will take %s to repay this loan!\n", formatPeriods(int(n)))
			fmt.Printf("Overpayment = %.0f\n", overpayment)

		case !provided["principal"]:
			pow := math.Pow(1+i, n)
			p = math.Floor(a / (i * pow / (pow - 1)))
			overpayment := n*a - p
			fmt.Printf("Your loan principal = %.0f!\n", p)
			fmt.Printf("Overpayment = %.0f\n", overpayment)
		}
	}
}
