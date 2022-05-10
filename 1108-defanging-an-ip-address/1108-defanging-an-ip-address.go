func defangIPaddr(address string) string {
    result := ""
    for _, ch := range address {
        if ch == '.' {
            result += "[.]"
            continue
        }
        result += string(ch)
    }
    return result
}