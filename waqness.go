package main

import "fmt"

type Car struct {
	Name, Model, Color string
	Year, Speed, HP    int64
}

func main() {
	c := Car{Name: "BMW", Model: "M5 e34", Color: "daytona violet", Year: 1995, Speed: 310, HP: 675}
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Model: ", c.Model)
	fmt.Println("Color: ", c.Color)
	fmt.Println("Year: ", c.Year)
	fmt.Println("Max Speed: ", c.Speed)
	fmt.Println("HP: ", c.HP)
}
