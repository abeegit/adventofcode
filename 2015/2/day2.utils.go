package main

import (
	"errors"
	"strconv"
	"strings"
)

func Sort(dimensions []int) []int {
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

func GetDimensions(dimensionString string) ([]int, error) {
	s := strings.Split(dimensionString, "x")
	if len(s) != 3 {
		return []int{}, errors.New("dimension string has invalid characters")
	}

	length, _ := strconv.Atoi(s[0])
	width, _ := strconv.Atoi(s[1])
	height, _ := strconv.Atoi(s[2])

	return []int{length, width, height}, nil
}

func CalculateVolume(dimensions []int) int {
	return dimensions[0] * dimensions[1] * dimensions[2]
}

