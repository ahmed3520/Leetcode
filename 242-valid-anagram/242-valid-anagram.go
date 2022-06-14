
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	store := make([]byte, 26)
	for i := range s {
		store[s[i]-'a']++
		store[t[i]-'a']--
	}

	for i := range store {
		if store[i] != 0 {
			return false
		}
	}
	return true

}

