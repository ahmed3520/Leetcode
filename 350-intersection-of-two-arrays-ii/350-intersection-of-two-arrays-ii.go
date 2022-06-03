func intersect(nums1 []int, nums2 []int) []int {
    itemsMap := make(map[int]int);
    min, max := nums1, nums2;
    if(len(nums1)> len(nums2)){
        min, max = nums2, nums1
    }
    newArray := make([]int, 0);
    
    for i :=0; i <len(min); i++ {
        itemsMap[min[i]]++;
    }
    fmt.Println(itemsMap)
    for j :=0; j < len(max);j++{
        _, ok := itemsMap[max[j]]
        if  ok{
            newArray = append(newArray, max[j])
            if itemsMap[max[j]]-1 == 0 {
				delete(itemsMap, max[j])
			} else {
				itemsMap[max[j]]--
			}
        }
        if len(itemsMap)==0 {
            return newArray
        }
    }
    return newArray
}