package offer

func levelprint29(nums [][]int) {
	l1 := 0
	l2 := len(nums[0])
	l3 := 0
	l4 := len(nums)
	var res []int
	for l1 <= l2 && l3 <= l4 {
		for i := l1; i < l2; i++ {
			res = append(res, nums[l3][i])
		}
		for i := l3; i < l4; i++ {
			res = append(res, nums[i][l2])
		}
		for i := l2 - 1; i >= l1; i-- {
			res = append(res, nums[l4][i])
		}
		for i := l4 - 1; i >= l2; i-- {
			res = append(res, nums[i][l1])
		}
		l1++
		l2--
		l3++
		l4--
	}
}
