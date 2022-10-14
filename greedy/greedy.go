package greedy

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	profit := 0
	for _, price := range prices {
		if price < buy {
			buy = price
		} else if (price - buy) > profit {
			profit = price - buy
		}

	}
	return profit
}

func longestPalindrome(s string) int {
	letters := make(map[rune]int)
	for _, letter := range s {
		letters[letter] += 1
	}
	odd, even := 0, 0
	for _, value := range letters {
		even += value / 2 * 2
		if value%2 != 0 {
			odd += 1
		}
	}
	if odd > 0 {
		even += 1
	}
	return even
}
