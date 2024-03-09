package main

import "fmt"


func main() {

	for x:=1; x<10; x++{
		fmt.Println(x)
	}

	y:=5

	for{
		fmt.Println("Just do it",y)
		y += 5
		if y>25 {
			break
		}
	}
}