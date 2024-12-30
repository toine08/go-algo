package algorithm

func Bubble_sort(nums []int) []int {
	swapping := true
	end := len(nums)
	for swapping {
		swapping = false
		for i := 1; i < end; i++ {
			if nums[i-1] > nums[i] {
				temp := nums[i-1]
				nums[i-1] = nums[i]
				nums[i] = temp
				swapping = true
			}
		}

		end--
	}

	return nums
}
