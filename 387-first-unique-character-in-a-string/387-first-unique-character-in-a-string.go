func firstUniqChar(s string) int {
    // two pass
    res := -1
    
    table := map[rune]int{}
    
    for _, ch:= range s{
        table[ch] += 1
    }
    
    for idx, ch:= range s{
        if table[ch]==1{
            res = idx
            break
        }
    }
    return res
    
}