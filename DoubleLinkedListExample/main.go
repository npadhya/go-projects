package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// Getthe length of both the slices
	m, n := len(nums1), len(nums2)

	// if first slice is longer than second, than swap both the slice and their lengths
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	// calculate initial min = 0 , max = length of the longer slice and mid = sum of lengths of both the slice + 1 . 
	iMin, iMax, halfLen := 0, m, (m+n+1)/2

	// While min is smaller than max
	for iMin <= iMax {
		// Get the avg of current min and max
		currentMid := (iMin + iMax) / 2
		// Get the initial mid point + 1
		updatedHalflan := halfLen - currentMid

		// if currentMid is smaller than the long slice
		// and 
		if currentMid < m && nums2[updatedHalflan-1] > nums1[currentMid] {
			iMin = currentMid + 1
		} else if currentMid > 0 && nums2[updatedHalflan] < nums1[currentMid-1] {
			iMax = currentMid - 1
		} else {
			maxOfLeft := 0
			if currentMid == 0 {
				maxOfLeft = nums2[updatedHalflan-1]
			} else if updatedHalflan == 0 {
				maxOfLeft = nums1[currentMid-1]
			} else {
				maxOfLeft = max(nums1[currentMid-1], nums2[updatedHalflan-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			minOfRight := 0
			if currentMid == m {
				minOfRight = nums2[updatedHalflan]
			} else if updatedHalflan == n {
				minOfRight = nums1[currentMid]
			} else {
				minOfRight = min(nums1[currentMid], nums2[updatedHalflan])
			}

			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println("Median:", median)
}