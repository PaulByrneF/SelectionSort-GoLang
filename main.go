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

	// Make copy of initial slice for selection sort algorithm
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Printf(" -- Selection Sort Algorithm --\n%v\n\n", selectionSort(slice2))

	// Make copy of initial slice for bubblesort algorithm
	slice3 := make([]int, len(slice))
	copy(slice3, slice)
	fmt.Printf(" -- Bubblesort Algorithm --\n%v\n\n", bubbleSort(slice3))
}

func bubbleSort(slice3 []int) []int {

	length := len(slice3)
	sorted := false

	for !sorted {
		swapped := false
		for ind := 0; ind < length-1; ind++ {
			if slice3[ind] > slice3[ind+1] {
				slice3[ind+1], slice3[ind] = slice3[ind], slice3[ind+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		length = length - 1
	}
	return slice3
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func selectionSort(slice []int) []int {
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
