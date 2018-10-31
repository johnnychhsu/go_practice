package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// receiver can be used driectly by pointer and value.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
    // check this, different from func
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
    // check this
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

