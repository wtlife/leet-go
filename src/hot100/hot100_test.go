package hot100

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums :=[]int{3,2,4}
	ans := TwoSum(nums,6)
	fmt.Println(ans)
}

func TestLengOfLS(t *testing.T)  {
	s := "abcabcbb"
	ans := lengthOfLongestSubstring(s)
	fmt.Print(ans)
}
func TestMedian(t *testing.T)  {
	nums1:= []int{1,2}
	nums2:= []int{3,4}
	ans:=findMedianSortedArrays(nums1,nums2)
	fmt.Print(ans)
}
func TestPalindrome(t *testing.T)  {
	s := "babad"
	ans := longestPalindrome(s)
	fmt.Print(ans)
}