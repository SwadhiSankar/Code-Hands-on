package main

import "fmt"

func main(){
	
    revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses :")
	taxRate := getUserInput("TaxRate : ")

	EBT, profit, ratio :=CalculateFinancials(revenue,expenses,taxRate)
	
	
	fmt.Printf("%.1f\n",EBT)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.3f\n",ratio)

}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64,float64,float64){

	EBT := revenue-expenses
	profit := EBT * (1-taxRate/100)
	ratio := EBT/profit
	return EBT,profit,ratio
}

func getUserInput(Text string)float64{
	var userInput float64
	fmt.Print(Text)
	fmt.Scan(&userInput)
	return userInput
}
