package main

import (
	"fmt"
)

func sort(array []int, leng int) {

	fmt.Println("sort方法中传进来的数值：", array)

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp

			}
		}

	}
	fmt.Println("sort方法中排序后的数组：", array)
}

func main() {

	array := []int{5, 7, 9, 4, 2, 1, 4, 15, 3, 18, 17}

	fmt.Println("原始的数组：", array)

	sort(array, len(array))

	fmt.Println("排序后的数组：", array)
}
