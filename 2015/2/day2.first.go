package main

import (
	"fmt"
	"strings"
)

func getAreaForDimensions(dimensionString string) (int, error) {
  dimensions, error := GetDimensions(dimensionString)
	if error != nil {
		return 0, error
	}	

	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	sortedDimensions := Sort([]int{length, width, height})

	areaOfSmallestSide := sortedDimensions[0] * sortedDimensions[1] 
	surfaceArea := 2*length*width + 2*width*height + 2*length*height
	
	area := surfaceArea + areaOfSmallestSide
	return area, nil
}


func First() {
	dimensionsSlice := strings.Split(input, "\n")

	total := 0
	for _, dimensionString := range dimensionsSlice {
		area, err := getAreaForDimensions(dimensionString)
		if err != nil {
			fmt.Println(err)
			break
		}
	
		total = total + area
	}

	fmt.Println("The total area of wrapping paper required is ", total)
}