package utils

// order and elements in the two slices have to match
func IdenticalSlices(s1 []interface{}, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
