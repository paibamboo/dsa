// Round up, Asc
// Result = right most
l, r := 0, len(nums) - 1
for l < r {
	m := l + (r-l+1)/2
	if target >= nums[m] {
		l = m
	} else {
		r = m - 1
	}
}

// Round up, Desc
// Result = right most
l, r := 0, len(nums) - 1
for l < r {
	m := l + (r-l+1)/2
	if target <= nums[m] {
		l = m
	} else {
		r = m - 1
	}
}

// Round down, Asc
// Result = left most
l, r := 0, len(nums) - 1
for l < r {
	m := l + (r-l)/2
	if target <= nums[m] {
		r = m
	} else {
		l = m + 1
	}
}

// Round down, Desc
// Result = left most
l, r := 0, len(nums) - 1
for l < r {
	m := l + (r-l)/2
	if target >= nums[m] {
		r = m
	} else {
		l = m + 1
	}
}

// Can reach -1th index and the nth index before it leaves the loop
// Useful for adding negative and positive infinity to the edge of a slice
// Example problem: https://leetcode.com/problems/median-of-two-sorted-arrays/description/
l, r := 0, len(nums)
for l <= r {
	m := (l + r)/2
	if m < 0 || m >= len(nums) {
		return -1 // Not found
	}
	if nums[m] == target {
		return m
	}
	if nums[m] < target {
		l = m + 1
	} else {
		r = m - 1
	}
	return -1 // Not found
}
