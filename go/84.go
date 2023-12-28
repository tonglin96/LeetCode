package _go

func largestRectangleArea(heights []int) int {
	var maxArea int
	for i := 0; i < len(heights); i++ {

		right := i + 1
		left := i - 1
		for left >= 0 && heights[left] > heights[i] {
			left++
		}
		for right < len(heights) && heights[right] > heights[i] {
			right++
		}
		maxArea = max(maxArea, (right-left+1)*heights[i])
	}
	return maxArea
}
