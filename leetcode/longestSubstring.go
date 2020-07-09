package leetcode

import (
	"strings"
)

//获取不重复的最长子串
func LongestSubstring(s string) (ss string) {
	t := ""
	for _, v := range s {
		idx := strings.Index(t, string(v))
		if idx >= 0 {
			if len(t) > len(ss) {
				ss = t
			}
			t = t[idx+1:]
		}
		t += string(v)
	}
	if len(t) > len(ss) {
		ss = t
	}
	return
}

func TestLongestSubstring() string {
	str := "abcabcdabcdeabc"
	return LongestSubstring(str)
}
