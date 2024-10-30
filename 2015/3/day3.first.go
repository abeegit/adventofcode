package main

import (
	"fmt"
	"unicode/utf8"
)

func First() {
	length := utf8.RuneCountInString(input)

	grid := make([][]int, length)
	
	for i := range grid {
		grid[i] = make([]int, length)
	}

	center := length/2

	homesDeliveredMap := make(map[string]int)
	homesDeliveredMap[GenerateKey(center, center)]++
	
	row := center
	col := center

	for _, move := range input {
		switch move {
			case '<':
				col--
			case '>': 
				col++
			case '^': 
				row--
			case 'v': 
				row++
		}

		houseKey := GenerateKey(row, col)
		homesDeliveredMap[houseKey]++
	}
	fmt.Println("Number of houses visited", len(homesDeliveredMap))
}