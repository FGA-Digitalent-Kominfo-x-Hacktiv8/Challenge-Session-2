package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {
		fmt.Println("Nilai j = ", j)
		if j == 5 {
			for idx, val := range "САШАРВО" {
				fmt.Printf("Character %U '%c' starts at byte position %d \n", val, val, idx)
			}
		}
	}

	// for j := 6; j <= 10; j++ {
	// 	fmt.Println("Nilai j = ", j)
	// }
}
