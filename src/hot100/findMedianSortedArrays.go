package hot100

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2

	val1:=findKth(nums1, 0, nums2, 0, left)
	val2:=findKth(nums1, 0, nums2, 0, right)
	return float64(val1+val2)/2
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k==1 {
		if nums1[i]<nums2[j] {
			return nums1[i]
		}else {
			return nums2[j]
		}
	}

	var midVal1, midVal2 int
	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = math.MaxInt32
	}

	if j+k/2-1 < len(nums1) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = math.MaxInt32
	}

	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}

