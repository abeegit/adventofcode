package main

import "fmt"

func GenerateKey(row int, col int) string {
	return fmt.Sprintf("%d::%d", row, col)
}

