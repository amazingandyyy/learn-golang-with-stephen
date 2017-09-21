package main

import "fmt"

func main() {
	// 1: var colors map[string]string
	// 2: colors := make(map[string]string)
	// 3:
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"green": "$4b745",
	}
	// > colors["white"] = "#ffffff" // use [] to access the value
	// > delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
