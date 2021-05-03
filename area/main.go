package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}
type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{
		base:   10.2,
		height: 30.8,
	}
	sq := square{
		side: 15.6,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(sh shape) {
	fmt.Println("The area valueis", sh.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return math.Pow(sq.side, 2)
}
