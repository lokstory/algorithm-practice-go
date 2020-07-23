package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Original way
func partition(list []int, left int, right int) int {
	pivot := list[right]

	// Put smaller than pivot to left
	for i := left; i <= right; i++ {
		if list[i] < pivot {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}

	// Swap pivot to smaller max id + 1 (center)
	list[left], list[right] = list[right], list[left]

	return left
}

func sort(list []int, low int, high int) []int {
	if low < high {
		pi := partition(list, low, high)

		// Sort left
		list = sort(list, low, pi-1)

		// Sort right
		list = sort(list, pi+1, high)
	}

	return list
}

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	left, right := 0, len(list)-1

	// Pick a pivot
	pivot := rand.Int() % len(list)

	// Move the pivot to the right
	list[pivot], list[right] = list[right], list[pivot]

	// Pile elements smaller than the pivot on the left
	for i := range list {
		if list[i] < list[right] {
			list[i], list[left] = list[left], list[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	list[left], list[right] = list[right], list[left]

	// Sort left
	quicksort(list[:left])

	// Sort right
	quicksort(list[left+1:])

	return list
}

func main() {
	values := []int{10, 7, 8, 9, 1, 5}

	now := time.Now()

	ret := quicksort(values)

	fmt.Println("ret:", ret)

	fmt.Println("duration:", time.Now().Sub(now))

	now = time.Now()

	ret2 := sort(values, 0, len(values)-1)

	fmt.Println("ret2:", ret2)

	fmt.Println("duration:", time.Now().Sub(now))
}
