package main

import (
	"fmt"
)

func second() {
	var level int = 0
	var basementReached bool = false

	for i := range input {
		step := input[i]	
		switch step {
			case '(':
				level++
			case ')':
				level--
		}

		if (!basementReached && level == -1) {
			basementReached = true
			fmt.Printf("Basement was reached at the character %d", i+1)
			break
		}
	}
}