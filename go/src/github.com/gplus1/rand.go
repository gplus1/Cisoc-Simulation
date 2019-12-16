package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand := randomint(1, 21)
	slice := make([]int, 5, 6)
	i := 0
	count := 1
	for i < 4 {
		var num int
		fmt.Println("Please enter your guess:  ")
		fmt.Scan(&num)
		if !repeatchk(slice, num) {
			i++
			count++
		}
		slice = append(slice, num)
		if num == rand {
			fmt.Println("You have won the game!")
			break
		} else if num > rand {
			fmt.Println("TooLarge")
		} else if num < rand {
			fmt.Println("TooSmall")
		}
	}
	fmt.Println(count)
	if count == 5 {
		fmt.Println("You have lost!")
	}

}

func randomint(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func repeatchk(slice []int, num int) bool {
	var guss bool = false
	for i := 0; i < len(slice); i++ {
		if num == slice[i] {
			guss := true
			return guss
			break
		}

	}
	return guss
}
