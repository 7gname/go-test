package xiaomi

//拼接两个有序数组
func ConcatArr(s1, s2 []int) (s []int) {
	i := 0
	j := 0
	for {
		if i >= len(s1) {
			s = append(s, s2[j:]...)
			break
		}
		if j >= len(s2) {
			s = append(s, s1[i:]...)
			break
		}

		if s1[i] < s2[j] {
			s = append(s, s1[i])
			i++
		} else {
			s = append(s, s2[j])
			j++
		}
	}
	return
}
