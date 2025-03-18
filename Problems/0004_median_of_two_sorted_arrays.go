func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var n1, n2 int = len(nums1), len(nums2)
    tot := n1 + n2
    point, len := tot / 2, 1
    if tot % 2 == 0 {
        len = 2
    }
    i, j, k := 0, 0, 0
    var cur, pre int
    for i < n1 && j < n2 && k <= point {
        if k > 0 {
            pre = cur
        }
        if nums1[i] < nums2[j] {
            cur = nums1[i]
            i++
        } else {
            cur = nums2[j]
            j++
        }
        k++
    }
    if k <= point {
        for i < n1 && k <= point {
            pre = cur
            cur = nums1[i]
            i++
            k++
        }
        for j < n2 && k <= point {
            pre = cur
            cur = nums2[j]
            j++
            k++
        }
    }
    if len ==1 {
        return float64(cur)
    } else {
        return float64(cur + pre) / 2.0
    }
}