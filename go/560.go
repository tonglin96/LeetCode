package _go

// 示例 1：
//
// 输入：nums = [1,1,1], k = 2
// 输出：2
// 示例 2：
//
// 输入：nums = [1,2,3,3], k = 3
// 输出：2
func subarraySum(nums []int, k int) int {
	var d = make(map[int]int)
	d[0] = 1
	var sum int
	var res int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := d[sum-k]; ok {
			res += d[sum-k]
		}
		d[sum]++
	}
	return res
}
