package main

import (
	"fmt"
	"strings"
)

func getLengthForRibbon(dimensionString string) (int, error) {
	dimensions, error := GetDimensions(dimensionString)
	if error != nil {
		return 0, error
	}	

	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	sortedDimensions := Sort([]int{length, width, height})

	smallestPerimeter := 2*(sortedDimensions[0] + sortedDimensions[1])
	volumeOfBox := length * width * height

	return smallestPerimeter + volumeOfBox, error
}

func Second() {
	dimensionsSlice := strings.Split(input, "\n")

	total := 0
	for _, dimensionString := range dimensionsSlice {
		length, err := getLengthForRibbon(dimensionString)
		if err != nil {
			fmt.Println(err)
			break
		}
	
		total = total + length
	}

	fmt.Println("Total length of ribbon required: ", total)
}