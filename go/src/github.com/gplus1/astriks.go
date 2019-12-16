package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 5
	for i := 5; i >= 0; i-- {
		fmt.Println(strings.Repeat(" ", i) + strings.Repeat("*", count-i) + strings.Repeat(" ", i))

	}
	count_2 := 5
	for i := 0; i <= 5; i++ {
		fmt.Println(strings.Repeat("*", count_2-i))

	}
}
func divmod(num int, dem int) int {
	if num > dem {
		return num / dem
	} else {
		return 1
	}

}
