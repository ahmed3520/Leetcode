
func maximumUniqueSubarray(nums []int) int {
	cur_sum, ans, l := 0, 0, 0
	s := map[int]bool{}
	for r, _ := range nums {
		for {
			if _, ok := s[nums[r]]; ok {
				delete(s, nums[l])
				cur_sum -= nums[l]
				l += 1
			} else {
				break

			}
		}
		s[nums[r]] = true

		cur_sum += nums[r]
		ans = max(ans, cur_sum)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
