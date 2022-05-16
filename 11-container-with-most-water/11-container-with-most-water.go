func maxArea(height []int) int {
    	max := 0
	i, j := 0, len(height)-1
	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}