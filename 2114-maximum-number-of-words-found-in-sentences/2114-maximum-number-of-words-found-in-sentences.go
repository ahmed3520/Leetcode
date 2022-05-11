func CountWords(s string) int {
	return len(strings.Fields(s))
}
func mostWordsFound(sentences []string) int {
    max :=0
    for _,ch := range sentences {
        temp:=CountWords(ch)
        if temp > max {
            max = temp
        }
    }
    return max
}