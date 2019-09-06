//TipCalc by Pj Chavez
package main

import (
 "fmt"
 "os"
 "os/exec"
)

//Clearscreen ...
func Clearscreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("cmd", "/c", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getbill() float64 {
	check := false
	var bill float64
	for check != true {
		fmt.Println("\nEnter the billed amount.")
		fmt.Scan(&bill)

	if bill >= .00 {
		check = true
	} else {
		Clearscreen()
		fmt.Println("Enter a valid amount.")
	}
}
	return bill
}

func gettip(b float64) float64 {
	check := false
	var tip float64
	for check != true {
		fmt.Println("\nEnter tip percentage.")
		fmt.Scan(&tip)

	if tip >= 0 && tip < 100 {
		check = true
	} else {
		Clearscreen()
		fmt.Println("Enter a valid amount.")
	}
}
	return (tip*.01)*b
}

func showfinalbill(b,t,tb float64) {
	Clearscreen()
	fmt.Print("TipCalc\n")
	fmt.Printf("Bill: $%.2f \n",b)
	fmt.Printf("Tip: $%.2f \n",t)
	fmt.Printf("Total Bill: $%.2f \n",b+t)
}

func main() {
	Clearscreen()
	fmt.Print("TipCalc\n")
	bill := 0.0
	tip := 0.0
	totalbill := 0.0
	
	bill = getbill()
	tip = gettip(bill)

	showfinalbill(bill,tip,totalbill)
}