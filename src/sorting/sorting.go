package sorting

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
}
