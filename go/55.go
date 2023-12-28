package _go

// 示例 1：
//
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例 2：
//
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
func canJump(nums []int) bool {
	var max_i int
	for i := 0; i < len(nums); i++ {
		if max_i >= i && i+nums[i] > max_i { // 能跳到目前的位置，且最远距离大于当前最远距离
			max_i = i + nums[i]
		}
	}
	return max_i >= len(nums)
}
