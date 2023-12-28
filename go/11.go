package _go

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var area int
	for left < right {
		if height[left] < height[right] {
			if area < (right-left)*height[left] {
				area = (right - left) * height[left]
			}
			left++
		} else {
			if area < (right-left)*height[right] {
				area = (right - left) * height[right]
			}
			right--
		}
	}
	return area
}
