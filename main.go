package main

import (
	"fmt"

	"github.com/toine08/go-algo/algorithm"
)


func main() {
	separator :=fmt.Sprintf("\n----------------------------------------------\n")
	nums := []int{1, 3, 2, 5, 4}
	nums_nonSorted := []int{1, 3, 2, 5, 4}
	nums_Duplicates := []int{1,3,2,2,4,4,5,2}

	fmt.Printf("Array before insertion sort: %v \n", nums)
	algorithm.Insertion_sort(nums)
	fmt.Printf("Array after insertion sort: %v \n", nums)
	fmt.Println(separator)

	nums = nums_nonSorted
	fmt.Printf("Array before bubble sort: %v \n", nums)
	algorithm.Bubble_sort(nums)
	fmt.Printf("Array after bubble sort: %v \n", nums)
	fmt.Println(separator)

	fmt.Printf("Array before removing duplicates: %v \n", nums_Duplicates)
	nums_Sorted := algorithm.Remove_duplicates(nums_Duplicates)
	fmt.Printf("Array sorted: %v \n", nums_Sorted)


}
