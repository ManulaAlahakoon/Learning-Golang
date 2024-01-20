package main

import "fmt"


//func add(x float64, y float64) float64 {

func add(x , y float64) float64 {
	return x + y
}

func multiple(x string, y string) (string,string) {
	return x,y
}

func main(){
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5

	//var num1,num2 float64 = 5.6,9.5 Mostly use outside of the funcs

	num1,num2 := 100.0,11.0// By default float64
	
	a,b := "I Love","You"

	fmt.Println(add(num1,num2))
	fmt.Println(multiple("Hello","Darling"))
	fmt.Println(multiple(a,b))

	//Type casting
	var z int = 10
	var w float32 = float32(z)
	p := w  //p will be type float32
	fmt.Println(p)


}


