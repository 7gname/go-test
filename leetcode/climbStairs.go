package leetcode

func ClimbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

func TestClimbStairs() int {
	return ClimbStairs(10)
}
