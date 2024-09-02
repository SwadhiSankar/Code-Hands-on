package main

import (
	"fmt"
	"math"
)

func main(){
	
	 const inflationRate = 2.5
	 var investmentAmount float64
	 var years, expectedReturnRate float64
	 fmt.Print("Enter the investment amount: ")
	 fmt.Scan(&investmentAmount)

	fmt.Print("Years")
	fmt.Scan(&years)

	fmt.Print("ExpectedRate")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100,years) 
	futureRealValue := futureValue/ math.Pow(1+inflationRate/100,years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}