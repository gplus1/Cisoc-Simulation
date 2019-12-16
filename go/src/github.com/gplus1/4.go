package main

import (
	"fmt"
)

func main() {
	list1 := []int{1, 2, 4}
	list2 := []string{"a", "b", "c"}

	list3 := list1

	for index, _ := range list2 {
		list3 = append(list3, list2[index])
	}
	fmt.Println(list3)
	fmt.Println(append(list1, list2...))
}
