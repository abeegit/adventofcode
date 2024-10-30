package main

import (
	"fmt"
	"unicode/utf8"
)

func Second() {
	length := utf8.RuneCountInString(input)

	grid := make([][]int, length)
	
	for i := range grid {
		grid[i] = make([]int, length)
	}

	center := length/2
	
	homesDeliveredMap := make(map[string]int)

	centerHomeKey := GenerateKey(center, center) 
	homesDeliveredMap[centerHomeKey]++
	
	santaRow, santaCol, roboSantaRow, roboSantaCol := center, center, center, center

	for i, move := range input {
		alteredI := i + 2

		isEven := alteredI % 2 == 0 

		var row, col int

		if isEven {
			row = santaRow
			col = santaCol
		} else {
			row = roboSantaRow
			col = roboSantaCol
		}

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
		
		if isEven {
			santaRow = row
			santaCol = col
		} else {
			roboSantaRow = row
			roboSantaCol = col
		}

		houseKey := GenerateKey(row, col)
		homesDeliveredMap[houseKey]++
	}
	fmt.Println("Number of houses visited", len(homesDeliveredMap))
}