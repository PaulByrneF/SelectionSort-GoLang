package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Initialize Unsorted Slice
	slice := generateSlice(24)
	fmt.Printf(" -- Unsorted --\n%v\n\n", slice)

	// Selection Sort Algorithm
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Printf(" -- Selection Sort Algorithm --\n%v\n\n", SelectionSort(slice2))

	// BubbleSort Algorithm
	slice3 := make([]int, len(slice))
	copy(slice3, slice)
	fmt.Printf(" -- Bubblesort Algorithm --\n%v\n\n", BubbleSort(slice3))

	// InsertionSort Algorithm
	slice4 := make([]int, len(slice))
	copy(slice4, slice)
	fmt.Printf(" -- Insertion Sort Algorithm --\n%v\n\n", InsertionSort(slice4))

}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func InsertionSort(slice []int) []int {
	length := len(slice)
	for outerInd := 0; outerInd < length; outerInd++ {
		innerInd := outerInd
		for innerInd > 0 {
			if slice[innerInd-1] > slice[innerInd] {
				slice[innerInd-1], slice[innerInd] = slice[innerInd], slice[innerInd-1]
			}
			innerInd = innerInd - 1
		}
	}
	return slice
}

func BubbleSort(slice []int) []int {

	length := len(slice)
	sorted := false

	for !sorted {
		swapped := false
		for ind := 0; ind < length-1; ind++ {
			if slice[ind] > slice[ind+1] {
				slice[ind+1], slice[ind] = slice[ind], slice[ind+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		length = length - 1
	}
	return slice
}

func SelectionSort(slice []int) []int {
	var length = len(slice)
	for outerInd := 0; outerInd < length; outerInd++ {
		var minIdx = outerInd
		for innerInd := outerInd; innerInd < length; innerInd++ {
			if slice[innerInd] < slice[minIdx] {
				minIdx = innerInd
			}
		}
		slice[outerInd], slice[minIdx] = slice[minIdx], slice[outerInd]
	}
	return slice
}
