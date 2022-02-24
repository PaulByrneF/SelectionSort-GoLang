package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(24)
	fmt.Printf(" -- Unsorted --\n%v\n\n", slice)
	fmt.Printf(" -- Selection Sorted --\n%v\n", selectionSort(slice))

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
	sliceCopy := slice[:]
	var length = len(sliceCopy)
	for outerInd := 0; outerInd < length; outerInd++ {
		var minIdx = outerInd
		for innerInd := outerInd; innerInd < length; innerInd++ {
			if sliceCopy[innerInd] < sliceCopy[minIdx] {
				minIdx = innerInd
			}
		}
		sliceCopy[outerInd], sliceCopy[minIdx] = sliceCopy[minIdx], sliceCopy[outerInd]
	}
	return slice
}