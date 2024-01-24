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

func (b *road) brick_price() float64 {
	b.bricks = 50
	return float64(b.bricks) * brick
}

func (b road) metal_price() int {
	return int(b.metal)*metal
}

// Pointer receiver method
func (p *road) metal_change(x uint16) {
	p.metal = x
}

//Alternative
func new_preasure(x road,p float64) road {
	x.preassur = p
	return x
}


func main(){
	road_a := road{metal: 50000,bricks: 650,carpet:4000,preassur:300.67}
	road_b := road{40000,30000,5000,67.898}
	fmt.Println(road_a.bricks)
	fmt.Println(road_b.preassur)

	fmt.Println(road_a.brick_price())

	road_b.metal_change(100)
	road_b = new_preasure(road_b,20.567)

	fmt.Println(road_b.preassur)

	fmt.Println(road_b.metal_price())
	fmt.Println(road_a.bricks)
}