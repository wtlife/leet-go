package hot100

func lengthOfLongestSubstring(s string) int {
	visit:=make(map[byte]bool)
	left :=0
	max:=0
	for right:=0;right<len(s);right++{
		cur := s[right]
		for visit[cur] && left<right{
			delete(visit, s[left])
			left++
		}

		visit[cur] = true
		if (right-left+1)>max{
			max = right-left+1
		}
	}
	return max
}
