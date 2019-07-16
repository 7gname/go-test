package leetcode

//func ThreeQuickSort(nums []int) ([]int) {
//	l := len(nums)
//	if l <= 1 {
//		return nums
//	}
//	flag := nums[0]
//	left := make([]int, 0)
//	middle := make([]int, 0)
//	right := make([]int, 0)
//	for _, v := range nums {
//		if v < flag {
//			left = append(left, v)
//		}else if v == flag {
//			middle = append(middle, v)
//		}else{
//			right = append(right, v)
//		}
//	}
//	rs := make([]int, l)
//	copy(rs[:len(left)], ThreeQuickSort(left))
//	copy(rs[len(left):len(left)+len(middle)], middle)
//	copy(rs[len(left)+len(middle):], ThreeQuickSort(right))
//
//	return rs
//}

func ThreeQuickSort(nums []int) ([]int) {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	flag := nums[0]
	t := -1
	i := 0
	n := l
	for i < l {
		if nums[i] < flag {
			nums[i], nums[t+1] = nums[t+1], nums[i]
			i++
			t++
		}else if nums[i] == flag {
			i++
		}else{
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n--
		}
	}
	return nums
}

func TestThreeQuicSort() ([]int) {
	nums := []int{2,0,2,1,1,0}
	return ThreeQuickSort(nums)
}
