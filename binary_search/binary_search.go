package binary_search

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	start, end := 0, len(nums)-1
	for {

		middle := (start + end) / 2

		if nums[middle] == target {
			return middle
		}

		if start == end {
			break
		}

		if target > nums[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}

		if start >= len(nums) {
			start = len(nums) - 1
		}

		if end < 0 {
			end = 0
		}

	}
	return -1

}

type VersionChecker struct {
	n int
}

func (c VersionChecker) isBad(version int) bool {
	if version >= c.n {
		return true
	} else {
		return false
	}
}

type versionChecker func(n int) bool

var isBadVersion versionChecker = VersionChecker{4}.isBad

func firstBadVersion(n int) int {
	start, end := 1, n
	result := n
	for start < end {
		middle := start + (end-start)/2
		if isBadVersion(middle) {
			end = middle
			result = middle
		} else {
			start = middle + 1

		}
	}
	return result
}
