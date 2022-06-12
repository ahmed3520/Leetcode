func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		v := nums[m]
        if v == target {
            return m
        }else if v < target{
			l = m + 1
        }else{
			r = m - 1
        }
	}
	return l
}