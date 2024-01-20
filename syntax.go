package main

import ("fmt"
		"math"
		"math/rand")

func cal(){
	fmt.Println("Square root of 4 is",math.Sqrt(4))
	fmt.Println("Random integer",rand.Intn(20))
}

func main(){
	cal()
}