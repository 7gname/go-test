package leetcode

//查找两数之和等于目标数
func TwoSum(nums []int, target int) [][]int {
	m := make(map[int]int)
	rs := make([][]int, 0)
	for i, v := range nums {
		complement := target - v
		if idx, exists := m[complement]; exists {
			rs = append(rs, []int{i, idx})
		}
		m[v] = i
	}
	return rs
}

func TestTowSum() [][]int {
	nums := []int{1, 2, 7, 11, 15}
	target := 9
	return TwoSum(nums, target)
}
