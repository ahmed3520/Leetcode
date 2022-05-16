func min(x, y int)int{
    if x <= y {
        return x
    }
    
    return y
}

func maxArea(height []int) int {
    maximum:=0
    l,r := 0, len(height)-1
    for l <r {
        area := (r-l) * min(height[l],height[r])
        if area > maximum {
            maximum = area
        }
        if height[l]<height[r]{
            l++
        }else {
            r--
        }
    }
    return maximum
}
