package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	hash := make(map[string]int)
	hash["quarter_of_a_dozen"] = 3
	hash["half_of_a_dozen"] = 6
	hash["dozen"] = 12
	hash["small_gross"] = 120
	hash["gross"] = 144
	hash["great_gross"] = 1728

	return hash
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if units[unit] == 0 {
		return false
	} else {
		if bill[item] != 0 {
			bill[item] += units[unit]
		} else {
			bill[item] = units[unit]
		}

		return true
	}
}
func main() {
	bill := NewBill()
	fmt.Println(bill)
}
