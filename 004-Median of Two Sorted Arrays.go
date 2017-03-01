package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	right := float64(findKth(nums1, nums2, (m+n)/2+1))
	if (m+n)%2 == 1 {
		return right
	}
	left := float64(findKth(nums1, nums2, (m+n)/2)) // some redundant work
	return float64(left+right) / 2.0
}

func findKth(nums1 []int, nums2 []int, k int) int { // k from 1 to m+n
	m, n := len(nums1), len(nums2)

	switch { // base cases
	case m == 0:
		return nums2[k-1]
	case n == 0:
		return nums1[k-1]
	case k == 1:
		if nums1[0] < nums2[0] {
			return nums1[0]
		}
		return nums2[0]
	}

	km, kn := (k-1)*(m-1)/(m+n), (k-1)*(n-1)/(m+n) // km + kn <= k - 2
	if nums1[km] < nums2[kn] {                     // len(0...km + 0...kn) <= k
		return findKth(nums1[(km+1):], nums2, k-km-1) // nums1[km] < Kth; km+1 <= m-1; k-km-1 >= 1;
	}
	return findKth(nums1, nums2[(kn+1):], k-kn-1) // nums2[kn] >= Kth; kn+1 <= n-1; k-kn-1 >= 1;
}
