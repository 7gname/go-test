package xiaomi

func SearchMin(s []int) (index int) {
	if len(s) == 1 {
		return 0
	}
	midpos := len(s)/2
	first := s[0]

	if first > s[midpos] {
		return midpos - SearchMin(s[0:midpos+1])
	}else{
		return midpos + SearchMin(s[midpos:])
	}
	return
}
