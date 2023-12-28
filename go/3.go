package _go

func lengthOfLongestSubstring(s string) int {
	var left, right int
	var existMap = make(map[byte]int)
	var maxLen int
	for right < len(s) {
		if _, ok := existMap[s[right]]; ok {
			delete(existMap, s[right])
			left++
		} else {
			if maxLen < right-left+1 {
				maxLen = right - left + 1
			}
			existMap[s[right]] = right
			right++
		}
	}
	return maxLen
}
