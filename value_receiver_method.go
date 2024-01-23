package main

import "fmt"

type road struct {
	metal uint16 //unsigned integer 16 bit - min 0 max 65536
	bricks uint16 
	carpet int16 // -32k +32k
	preassur float64
}

const brick float64 = 50
const metal int = 10

func (b road) brick_price() float64 {
	return float64(b.bricks) * brick
}

func (b road) metal_price() int {
	return int(b.metal)*metal
}


func main(){
	road_a := road{metal: 50000,bricks: 650,carpet:4000,preassur:300.67}
	road_b := road{40000,30000,5000,67.898}
	fmt.Println(road_a.bricks)
	fmt.Println(road_b.preassur)
	fmt.Println(road_a.brick_price())
	fmt.Println(road_b.metal_price())
}