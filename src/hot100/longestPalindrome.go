package hot100

func longestPalindrome(s string) string {
	var ans = ""
	for idx := 0; idx < len(s); idx++ {
		leftStr := expand(s, idx, idx)
		rightStr := expand(s, idx, idx+1)

		var tmp string
		if len(leftStr) > len(rightStr) {
			tmp = leftStr
		}else {
			tmp = rightStr
		}

		if len(tmp) > len(ans) {
			ans = tmp
		}
	}
	return ans
}

func expand(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
