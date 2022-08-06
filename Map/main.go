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

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existsItem := units[unit] //boolean
	_, existsBill := bill[item]
	if !existsBill || !existsItem {
		return false
	}
	quant := bill[item]
	if quant < 0 {
		return false
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
func main() {
	bill := NewBill()
	fmt.Println(bill)
}
