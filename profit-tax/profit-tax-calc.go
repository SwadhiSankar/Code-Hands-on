package main

import "fmt"

func main(){
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue")
	fmt.Scan(&revenue)

	fmt.Print("expenses")
	fmt.Scan(&expenses)

	fmt.Print("taxRate")
	fmt.Scan(&taxRate)

	EBT := revenue-expenses
	profit := EBT * (1-taxRate/100)
	ratio := EBT/profit

	formattedString := fmt.Sprintf("Profit :  %v\n",profit)
	
	// fmt.Printf("Profit :  %v\n",profit)
	fmt.Print(formattedString)
	fmt.Print(EBT)
	fmt.Print(ratio)




}