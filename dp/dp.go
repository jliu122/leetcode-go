package dp

func longestPalindrome(s string) string {

	length := len(s)
	if length < 2 {
		return s
	}
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	result := ""
	ml := 0
	for j := 0; j < length; j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = true
			} else if i == j - 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i + 1][j - 1]
			}
			l := j - i + 1
			if dp[i][j] && l > ml {
				ml = l
				result = s[i : j + 1]
			}
		}
	}
	return result
}
