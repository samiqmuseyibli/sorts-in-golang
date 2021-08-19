package main

import "fmt"

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}

func main() {
	arr := []int{2, 1, 5, 6, 7, 9}
	arr = selectionSort(arr)
	fmt.Println(arr)
}
