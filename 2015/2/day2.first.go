package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// create a total variable that counts the SA and excess
// split each dimension by the "x" sep
// spread the elements onto different variables for l w h
// calculate surface area of the box and add to total
// find the smallest two of the 3 dimensions
// add their product to the SA total

func getAreaForDimensions(dimensionString string) (int, error) {
	s := strings.Split(dimensionString, "x")
	if len(s) != 0 {
		return 0, errors.New("Dimension string has invalid characters")
	}

	length, _ := strconv.Atoi(s[0])
	width, _ := strconv.Atoi(s[1])
	height, _ := strconv.Atoi(s[2])

	var areaOfSmallestSide int
	switch {
	case length < width && width < height:
		areaOfSmallestSide = length * width 
	case length < width && height < width:
		areaOfSmallestSide = length * height
	case width < length && height < length:
		areaOfSmallestSide = width * height
	}

	surfaceArea := 2*length*width + 2*width*height + 2*length*height
	return surfaceArea + areaOfSmallestSide, nil
}


func First() {
	dimensionsSlice := strings.Split(input, "\n")
	fmt.Println("dimensions size is", utf8.RuneCountInString(dimensionsSlice[1]), len(dimensionsSlice))

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