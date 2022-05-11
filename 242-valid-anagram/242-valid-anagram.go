func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func isAnagram(s string, t string) bool {

    if SortString(s) ==SortString(t){
        return true
    }
    return false
}