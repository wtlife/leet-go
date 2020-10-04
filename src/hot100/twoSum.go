package hot100

func TwoSum(nums []int, target int) []int {

	idxMap := make(map[int]int)
	for idx, v := range nums {
		if _, ok := idxMap[target-v]; ok {
			ans := []int{idxMap[target-v], idx}
			return ans
		}
		idxMap[v] = idx
	}

	return nil
}
