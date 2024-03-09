package main

import "fmt"

func main(){
	Grades := make(map[string]float32)

	Grades["Ravishan"] = 82
	Grades["Manula"] = 78
	Grades["Lara"] = 91

	LarasGrade := Grades["Lara"]
	fmt.Println(LarasGrade)

	delete(Grades,"Lara")
	fmt.Println(Grades)

	for k,v := range Grades {
		fmt.Println(k,":",v)
	}
}