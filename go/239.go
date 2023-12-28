package _go

// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]

func maxSlidingWindow(nums []int, k int) []int {
	//单调队列 + 双指针
	queue := []int{}
	var res []int
	for right := 0; right < len(nums); right++ {
		//队列不空且考察元素大于队尾元素 移除队尾元素
		//要么队列为空 ， 要么考察元素小于队尾元素
		for len(queue) > 0 && nums[right] >= nums[queue[len(queue)-1]] {
			//queue 放入的是index 下标
			queue = queue[:len(queue)-1]
		}
		//这时候考察元素一定小于队尾元素
		queue = append(queue, right)

		left := right - k + 1
		//queue 存放的是下标 left 是左极限 滑动要将数组弹出
		for queue[0] < left {
			queue = queue[1:]
		}
		//right = k-1 时候 说明第一个窗口形成 那么可以加入res
		//之后right 每次++ 都会放入一个res
		if right >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
