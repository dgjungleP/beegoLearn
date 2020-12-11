package main

import "fmt"

func predictPartyVictory(senate string) string {
	senateArr := []byte(senate)
	loop := true
	for loop {
		rCount := 0
		dCount := 0
		for index, chartV := range senateArr {
			if chartV == 'R' {
				rCount++
				dCount = helper(index, senateArr, 'D', dCount)
			} else if chartV == 'D' {
				dCount++
				rCount = helper(index, senateArr, 'R', rCount)
			}
		}
		if rCount < 0 {
			return "Dire"
		} else if dCount < 0 {
			return "Radiant"
		}
	}
	return ""
}
func helper(index int, senateArr []byte, str byte, count int) int {
	for i, chart := range senateArr[index:] {
		if chart == str {
			senateArr[index+i] = ' '
			return count
		}
	}
	if count > 0 {
		helper(0, senateArr, str, 0)
	}
	return count - 1
}

func main() {
	fmt.Println(predictPartyVictory("RDDDRDRDD"))
}
