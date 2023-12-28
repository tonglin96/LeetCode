package _go

func trap(height []int) int {
	var leftList []int
	var maxLeft, maxRight int
	var rightList = make([]int, len(height))
	for i := 0; i < len(height); i++ {
		maxLeft = max(maxLeft, height[i])
		leftList = append(leftList, maxLeft)
	}

	for i := len(height) - 1; i >= 0; i-- {
		maxRight = max(maxRight, height[i])
		rightList[i] = maxRight
	}
	var area int
	for i := 0; i < len(height); i++ {
		if leftList[i] > rightList[i] && rightList[i] > height[i] {
			area += rightList[i] - height[i]
		} else if leftList[i] <= rightList[i] && leftList[i] > height[i] {
			area += leftList[i] - height[i]
		}
	}
	return area
}

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}
