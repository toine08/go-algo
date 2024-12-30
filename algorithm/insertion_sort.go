package algorithm

// nice to sort quickly very small dataset or partially sorted dataset, stable
func Insertion_sort(nums []int) []int {
	for i := range nums {
		j := i
		for j > 0 && nums[j-1] > nums[j] {
			temp := nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = temp
		}

	}
	return nums
}
