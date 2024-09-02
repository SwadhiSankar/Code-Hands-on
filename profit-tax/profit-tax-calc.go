package main

import (
	"errors"
	"fmt"
	"os"
)


func main(){
	
    revenue,err1 := getUserInput("Revenue: ")
	
	expenses,err2 := getUserInput("Expenses :")
	
	taxRate ,err3 := getUserInput("TaxRate : ")
	if err1 != nil || err2 !=nil || err3 != nil{
		fmt.Println(err1)
		return
	}

	EBT, profit, ratio :=CalculateFinancials(revenue,expenses,taxRate)
	
	
	fmt.Printf("%.1f\n",EBT)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.3f\n",ratio)
	storeResults(EBT,profit,ratio)

}

func storeResults (EBT, profit,ratio float64 ){
	results := fmt.Sprintf("EBT : %.1f\n Profit :%.1f \n Ratio: %.3f \n",EBT, profit,ratio)
	os.WriteFile("results.txt",[]byte (results), 0644)
}
func CalculateFinancials(revenue, expenses, taxRate float64) (float64,float64,float64){

	EBT := revenue-expenses
	profit := EBT * (1-taxRate/100)
	ratio := EBT/profit
	return EBT,profit,ratio
}

func getUserInput(Text string)(float64,error){
	var userInput float64
	fmt.Print(Text)
	fmt.Scan(&userInput)
	if userInput <=0{
		return 0,errors.New("Value should not be less than 0")
	}
	return userInput, nil
}