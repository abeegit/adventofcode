package main

import (
	"fmt"
	"unicode/utf8"
)

func generateKey(row int, col int) string {
	return fmt.Sprintf("%d::%d", row, col)
}

func First() {
	length := utf8.RuneCountInString(input)

	grid := make([][]int, length)
	
	for i := range grid {
		grid[i] = make([]int, length)
	}

	center := length/2
	fmt.Println("center is", center)

	homesDeliveredMap := make(map[string]int)
	homesDeliveredMap[generateKey(center, center)]++
	
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

		houseKey := generateKey(row, col)
		homesDeliveredMap[houseKey]++
	}
	fmt.Println("Number of houses visited", len(homesDeliveredMap))
}