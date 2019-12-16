package main

import (
	"fmt"
)

func main() {
	fmt.Println("Multiplication Table /n")

	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {

			fmt.Println("|    ", i, "   |  * ", "  | ", j, "  | ", i*j, "  |  ")

		}
		fmt.Println("|---------|--------|---------|")

	}

}
