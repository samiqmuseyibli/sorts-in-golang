package main

import "fmt"

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 1, 5, 6, 7, 9}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}
