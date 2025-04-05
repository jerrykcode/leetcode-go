func search(nums []int, target int) int {
	if target == nums[0] {
		return 0
	}
	l, r, mid, pivot := 0, len(nums)-1, 0, 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] >= nums[0] {
			pivot = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if target > nums[0] {
		return binary(nums, 0, pivot, target)
	} else {
		return binary(nums, pivot+1, len(nums)-1, target)
	}
}
func binary(nums []int, l int, r int, target int) int {
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}