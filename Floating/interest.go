package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0.0 {
		return 3.213
	} else if balance >= 0.0 && balance < 1000.0 {
		return 0.5
	} else if balance >= 1000.0 && balance < 5000.0 {
		return 1.621
	} else {
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	percent := float64(InterestRate(balance)) / 100.0
	return percent * balance

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	//goal := targetBalance - balance
	years := 0
	for balance < targetBalance {
		tax := float64(Interest(balance))
		balance += tax
		years += 1
	}
	return years
}
