package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit, depositRound float64
	var percentagePerMonth float64
	var yearsInvestPeriod int
	var bankIncome, diffDeposit float64

	for {
		fmt.Print("Введите сумму депозита: ")
		fmt.Scan(&depositRound)
		if depositRound <= 0 {
			fmt.Printf("Введите положительное чиcло типа Float64\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите процентную ставку в месяц: ")
		fmt.Scan(&percentagePerMonth)
		if percentagePerMonth <= 0 {
			fmt.Printf("Введите положительное чиcло типа Float64\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите количество лет: ")
		fmt.Scan(&yearsInvestPeriod)
		if yearsInvestPeriod <= 0 {
			fmt.Printf("Введите чиcло в диапозоне от 0 до %v\n", math.MaxInt)
		} else {
			break
		}
	}

	monthsInvestPeriod := yearsInvestPeriod * 12
	percentagePerMonth /= 100

	for i := 1; i <= monthsInvestPeriod; i++ {
		deposit = depositRound + depositRound*percentagePerMonth
		depositRound = math.Trunc(deposit*100) / 100
		diffDeposit = deposit - depositRound
		bankIncome += diffDeposit
	}

	fmt.Printf("Ваш депозит: %f, доход банка за счет округления: %v", depositRound, bankIncome)

}
