package string

func isIsomorphic(s string, t string) bool {
	characters_s := make(map[byte]byte)
	characters_t := make(map[byte]byte)
	for i := 0; i < len(s); i++ {

		if _, ok := characters_s[s[i]]; !ok {
			if _, ok := characters_t[t[i]]; ok {
				return false
			}
			characters_t[t[i]] = 0
			characters_s[s[i]] = t[i]
		}

		if characters_s[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, y := 0, 0; i < len(t); i++ {
		if (len(s) - y) > (len(t) - i) {
			return false
		}

		if s[y] == t[i] {
			y += 1
		}

		if y == len(s) {
			return true
		}

	}
	return false
}
