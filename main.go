package main

import "fmt"

//nice to sort quickly very small dataset or partially sorted dataset, stable
func insertion_sort(nums []int) []int{
	for i := range nums{
		j := i
		for j > 0 && nums[j-1] > nums[j]{
			temp:= nums[j]
			nums[j] = nums[j-1]
			nums[j-1] =temp
		}

	}
	return nums
}



func main(){
	nums := []int{1,3,2,5,4}
	fmt.Printf("Array before insertion_sort: %v \n", nums)
	insertion_sort(nums)
	fmt.Printf("Array after insertion_sort: %v \n", nums)

}