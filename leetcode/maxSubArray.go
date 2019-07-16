package leetcode

//求子数组最大之和
func MaxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		t := nums[i-1] + nums[i]
		if t > nums[i] {
			nums[i] = t
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func TestMaxSubArray() int {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	return MaxSubArray(nums)
}
