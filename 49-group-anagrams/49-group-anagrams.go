func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	mapString := make(map[string][]string)

	for _, w := range strs {
		word := SortString(w)
	    mapString[word] = append(mapString[word], w)
	}
    fmt.Println(mapString)
	var ss [][]string
	for _, value := range mapString {
		ss = append(ss, value)
	}
	return ss
}
