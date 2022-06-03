func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

    setCount := 0
	set := make(map[int]uint8)
    
	for _, n := range nums1 {
		set[n]++
        setCount++
	}

	ret := make([]int, 0, len(nums1))

	for _, n := range nums2 {
		if count := set[n]; count > 0 {
			ret = append(ret, n)
			set[n]--
            setCount--
		}
        if setCount == 0 {
            break
        }
	}

	return ret
}

