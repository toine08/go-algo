package main

import (
	"fmt"

	"github.com/toine08/go-algo/algorithm"
)

//nice to sort quickly very small dataset or partially sorted dataset, stable

func main() {
	nums := []int{1, 3, 2, 5, 4}
	nums_nonSorted := []int{1, 3, 2, 5, 4}

	fmt.Printf("Array before insertion sort: %v \n", nums)
	algorithm.Insertion_sort(nums)
	fmt.Printf("Array after insertion sort: %v \n", nums)
	fmt.Println("\n----------------------------------------------\n")
	nums = nums_nonSorted
	fmt.Printf("Array before bubble sort: %v \n", nums)
	algorithm.Bubble_sort(nums)
	fmt.Printf("Array after bubble sort: %v \n", nums)

}
