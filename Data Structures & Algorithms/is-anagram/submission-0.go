func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }
	s1 := make(map[byte]int, len(s))
	s2 := make(map[byte]int, len(s))
	for i := range len(s) {
		s1[s[i]] += 1
		s2[t[i]] += 1
	}
	for k, v := range s1 {
		if v != s2[k] {
			return false
		}
	}
	return true
}
