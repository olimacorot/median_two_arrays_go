package median_two_arrays

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    nums := append(nums1, nums2...)
    sort.Ints(nums[:])
    half := len(nums) / 2
    var total float64
   
    if len(nums) % 2 == 0 {
        total = (float64(nums[half-1]) + float64(nums[half])) / 2
    } else {
        total = float64(nums[half])
    }

    return total
}