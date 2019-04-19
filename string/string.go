package string

import "strings"

// 14
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs {
		for !strings.HasPrefix(s, result) {
			result = result[:len(result) - 1]
		}
	}
	return result
}
