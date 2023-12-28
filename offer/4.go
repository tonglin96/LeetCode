package offer

func search4(nums [][]int, target int) bool {
	level := 0
	idx := len(nums[0])
	for level < len(nums) && idx >= 0 {
		if nums[level][idx] == target {
			return true
		} else if nums[level][idx] < target {
			level++
		} else {
			idx--
		}
	}
	return false
}
