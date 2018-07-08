package stringset

func Contains(set []string, item string) bool {
	for _, element := range set {
		if element == item {
			return true
		}
	}
	return false
}

func With(set []string, items ...string) []string {
	result := []string{}
	for _, element := range append(set, items...) {
		if !Contains(result, element) {
			result = append(result, element)
		}
	}
	return result
}
