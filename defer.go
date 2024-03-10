package main

import "fmt"

func do() {
	defer fmt.Println("1) Defer is used to ensure that a function call is performed later in a program's execution.")
	defer fmt.Println("2) Defer print up-side-down")
	fmt.Println("What is defer?")
}

func repeat() {
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	do()
	repeat()
}