package main

import "fmt"

type Bill struct {
	receiptStr string
	rent       Rent
}

func (bill *Bill) PayBill(rent Rent) bool {
	amount := bill.generateBill(rent)
	fmt.Println("Bill Paid, Rs.", amount)
	return true
}

func (bill *Bill) generateBill(rent Rent) float64 {
	bill.rent = rent
	// logic to generate bill
	return 100.0
}
