package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Method attached to Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method that modifies the Rectangle
func (r *Rectangle) Resize(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("Area:", rect.Area())

	rect.Resize(2) // Double the size
	fmt.Printf("After resize: %+v\n", rect)
	fmt.Println("New area:", rect.Area())
}
