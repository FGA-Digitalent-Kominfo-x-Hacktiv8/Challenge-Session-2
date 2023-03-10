package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 4; j++ {
		fmt.Println("Nilai j = ", j)
	}

	for idx, val := range "САШАРВО" {
		fmt.Printf("Character %U starts at byte position %d \n", val, idx)
	}

	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j = ", j)
	}
}
