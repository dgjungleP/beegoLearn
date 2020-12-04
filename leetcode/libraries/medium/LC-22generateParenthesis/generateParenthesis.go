package main

func generateParenthesis(n int) []string {
	result := []string{}
	var dfs func(left, right int, result string)
	dfs = func(left, right int, path string) {
		if 2*n == len(path) {
			result = append(result, path)
			return
		}
		if left > 0 {
			dfs(right-1, right, path+"(")
		}
		if left < right {
			dfs(left, right-1, path+")")
		}
	}
	dfs(n, n, "")
	return result
}

func main() {

}
