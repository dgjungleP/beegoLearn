package main

func findCircleNum(isConnected [][]int) int {
	cityMap := make(map[int]bool)
	count := 0
	var dfs func(int, bool)
	dfs = func(city int, new bool) {
		if _, ok := cityMap[city]; !ok {
			cityMap[city] = true
			if new {
				count++
			}
			for number, child := range isConnected[city] {
				if _, ok := cityMap[number]; !ok && child == 1 {
					dfs(number, false)
				}
			}
		}
	}
	for i := 0; i < len(isConnected); i++ {
		dfs(i, true)
	}
	return count
}

func main() {
	findCircleNum([][]int{[]int{1, 0, 0, 1}, []int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 0, 1, 1}})
}
