

func canConstruct(ransomNote string, magazine string) bool {
    
    have := [26]int{}
    
    for _, char := range magazine {
        have[char-'a']++
    }
    
    for _, char := range ransomNote {
        have[char-'a']--
        if have[char-'a'] < 0 {
            return false
        }
    }
    
    return true
    
}

