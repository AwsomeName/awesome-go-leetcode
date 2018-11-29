package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	midian := make([]int, len(nums1)+len(nums2))
	m, n := 0, 0
	for i := 0; i < len(midian)/2+1; i++ {
		if m >= len(nums1) {
			midian[i] = nums2[n]
			n++
		} else if n >= len(nums2) {
			midian[i] = nums1[m]
			m++
		} else {
			if nums1[m] <= nums2[n] {
				midian[i] = nums1[m]
				m++
			} else {
				midian[i] = nums2[n]
				n++
			}
		}
	}
	if len(midian)%2 == 0 {
		return float64(midian[len(midian)/2]+midian[len(midian)/2-1]) / 2
	} else {
		return float64(midian[(len(midian)-1)/2])
	}
}
