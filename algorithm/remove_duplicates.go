package algorithm
import "slices"

func Remove_duplicates(nums []int) []int{
	newArr := []int{}

	for i := range nums{
		if !slices.Contains(newArr, nums[i]){
			newArr = append(newArr,nums[i])
		}
	}
	return newArr
}