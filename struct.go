package main

import "fmt"

type road struct {
	metal uint16 //unsigned integer 16 bit - min 0 max 65536
	bricks uint16 
	carpet int16 // -32k +32k
	preassur float64
}

func main(){
	road_a := road{metal: 50000,bricks: 65000,carpet:4000,preassur:300.67}
	road_b := road{40000,30000,5000,67.898}
	fmt.Println(road_a.bricks)
	fmt.Println(road_b.preassur)
}