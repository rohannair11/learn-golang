package main

import "fmt"

func main() {
	var ans_count uint64
	fmt.Println("Quiz Game")
	var name string
	fmt.Println("Enter your name = ")
	fmt.Scan(&name)
	var age uint64
	fmt.Println("Enter your age = ")
	fmt.Scan(&age)
	if (age >= 10){
		fmt.Println("you can proceed")
	} else {
		fmt.Println("lol")
	}

	fmt.Println("Lets begin the game ")
	fmt.Println("Capital Of India = Delhi Or Mumbai")
	var ans_capital string
	fmt.Scan(&ans_capital)

	if (ans_capital == "Delhi"){
		fmt.Println("Correct Answer")
		ans_count += 1
	} else {
		fmt.Println(" nah ")
	}

	fmt.Println("Business Capital Of India = UP Or Mumbai")
	var ans_bus_capital string
	fmt.Scan(&ans_bus_capital)

	if (ans_bus_capital == "Mumbai"){
		fmt.Println("Correct Answer")
		ans_count += 1
		} else {
		fmt.Println(" nah ")
	}
	fmt.Println("points = ", ans_count)




}
