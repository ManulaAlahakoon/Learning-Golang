package main

import "fmt"

func main(){

	x := 27
	address := &x
	fmt.Println(address)
	fmt.Println(*address)
	*address = 108
	fmt.Println(x)
	x = *address + 1
	fmt.Println(*address)
	*address = *address**address
	fmt.Println(x)

}