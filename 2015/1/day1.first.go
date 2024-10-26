package main

import "fmt"


func first() {
	var level int = 0

	for i := range input {
		step := input[i]	
		switch step {
			case '(':
				level++
			case ')':
				level--
			default:
				break
		}
	}

	fmt.Printf("Santa should go to the %d floor", level)
}
