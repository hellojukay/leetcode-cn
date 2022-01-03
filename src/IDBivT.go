
func generateParenthesis(n int) []string {
	var result []string
	dfs("", 0, 0, n, &result)
	return result
}

func dfs(current string, left int, right int, n int, result *[]string) {
	if left == n && right == n {
		*result = append(*result, current)
		return
	}
	if right > left {
		return
	}
	if left > n {
		return
	}
	dfs(current+"(", left+1, right, n, result)
	dfs(current+")", left, right+1, n, result)
}