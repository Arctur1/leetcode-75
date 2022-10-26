package stack

func backspaceCompare(s string, t string) bool {

	return applyRule(s) == applyRule(t)
}

func applyRule(s string) string {
	stack := []rune{}

	for _, c := range s {
		if c == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

// Input: s = "3[a2[c]]"
// Output: "accaccacc"
func decodeString(s string) string {
	return ""
}
