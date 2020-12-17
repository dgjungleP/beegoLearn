package main

func maxProfit(prices []int, fee int) int {
	sell, pay := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = max(sell, pay+prices[i]-fee)
		pay = max(pay, sell-prices[i])
	}
	return sell
}

func max(fir, src int) int {
	if fir > src {
		return fir
	}
	return src
}
func main() {

}
