package main

import "fmt"

func main() {
	// 1: var colors map[string]string
	// 2: colors := make(map[string]string)
	// 3:

	// > colors["white"] = "#ffffff" // use [] to access the value
	// > delete(colors, "white")
	myColorsList.printMap()
}

type colors map[string]string

func (c colors) printMap() {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

var myColorsList = colors{
	"red":   "#ff0000",
	"black": "#000000",
	"green": "$4b745",
}
