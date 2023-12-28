package _go

//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//示例 3:
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//func minWindow(s string, t string) string {
//	left := 0
//	right := 0
//	var keyMap = make(map[rune]int)
//	for _, i := range t {
//		if _, ok := keyMap[i]; ok {
//			keyMap[i] += 1
//		} else {
//			keyMap[i] = 1
//		}
//	}
//	var searchMap = make(map[rune]int)
//	for right < len(s) && left <= right {
//		if searchMap != keyMap {
//
//		}
//	}
//}
