package main

import (
	"fmt"
)

func main() {

	array := [12]int{5, 7, 9, 4, 2, 1, 4, 15, 3, 18, 17}

	fmt.Println("in order %d \n", array)
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp

			}
		}
	}

	fmt.Println("from vvvvv high to low %d \n", array)
}
