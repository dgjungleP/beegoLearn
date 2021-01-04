package main

func letterCombinations(digits string) []string {
	dir := make(map[string][]string)
	dir["2"] = []string{"a", "b", "c"}
	dir["3"] = []string{"d", "e", "f"}
	dir["4"] = []string{"g", "h", "i"}
	dir["5"] = []string{"j", "k", "l"}
	dir["6"] = []string{"m", "n", "o"}
	dir["7"] = []string{"p", "q", "r", "s"}
	dir["8"] = []string{"t", "u", "v"}
	dir["9"] = []string{"w", "x", "y", "z"}
	result := make([]string, 0)
	for _, c := range digits {
		if len(result) == 0 {
			if value, ok := dir[string(c)]; ok {
				result = append(result, value...)
			}
		} else {
			temp := make([]string, 0)
			for _, val := range result {
				if value, ok := dir[string(c)]; ok {
					for _, v := range value {
						temp = append(temp, val+v)
					}
				}

			}
			result = temp
		}
	}
	return result
}

func main() {
	letterCombinations("123")
}
