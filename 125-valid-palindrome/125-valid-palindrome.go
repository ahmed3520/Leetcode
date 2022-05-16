func strip(s string) string {
    var result strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        if ('a' <= b && b <= 'z') ||
            ('0' <= b && b <= '9') ||
            b == ' ' {
            result.WriteByte(b)
        }
    }
    return result.String()
}

func isPalindrome(s string) bool {
    newS := strings.ToLower(s)
    afterRmovingNonALpha := strip(newS)
    newTrim := strings.ReplaceAll(afterRmovingNonALpha, " ", "")
    sLen := len(newTrim)
    fmt.Println(newTrim)
    for i,j := 0,sLen-1; i < sLen ;i++ {
        if newTrim[i] != newTrim[j]{
            fmt.Println(newTrim[i], newTrim[j])
            
            return false
        }
        j--
    }
    return true
}
