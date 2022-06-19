package main

import "fmt"

type rect struct {
	x, y int
}

func (r rect) area() int {
	return r.x * r.y
}

func (r *rect) perim() int {
	return 2*r.x + 2*r.y
}

func main() {
	rec := rect{x: 10, y: 20}
	fmt.Println(rec.area())
	fmt.Println(rec.perim())
}