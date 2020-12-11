package main

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
			break
		case 10:
			five--
			ten++
			break
		case 20:
			if ten > 0 {
				ten--
			} else {
				five = five - 2
			}
			five--
			break
		}
		if ten < 0 || five < 0 {
			return false
		}
	}
	return true
}

func main() {
	lemonadeChange([]int{})
}
