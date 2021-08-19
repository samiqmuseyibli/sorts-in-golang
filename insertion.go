package main

import "fmt"

func insertionSort(arr []int) []int {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] >= temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}

func main() {
	arr := []int{2, 1, 5, 6, 7, 9}
	arr = insertionSort(arr)
	fmt.Println(arr)
}
