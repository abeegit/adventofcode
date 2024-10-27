package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func sort(dimensions []int) []int {
	for i := 0; i < 2; i++ {
		for j := i+1; j <= 2; j++ {
			if (dimensions[i] > dimensions[j]) {
				temp := dimensions[i]
				dimensions[i] = dimensions[j]
				dimensions[j] = temp
			}
		}
	}

	return dimensions
}

func getAreaForDimensions(dimensionString string) (int, error) {
	s := strings.Split(dimensionString, "x")
	if len(s) != 3 {
		return 0, errors.New("dimension string has invalid characters")
	}

	length, _ := strconv.Atoi(s[0])
	width, _ := strconv.Atoi(s[1])
	height, _ := strconv.Atoi(s[2])

	sortedDimensions := sort([]int{length, width, height})

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