package _go

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(curr_str string, left, right int)
	dfs = func(curr_str string, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, curr_str)
			return
		}
		if right < left {
			return
		}
		if left > 0 {
			dfs(curr_str+"(", left-1, right)
		}
		if right > 0 {
			dfs(curr_str+")", left, right-1)
		}

	}
	dfs("", n, n)
	return res
}
