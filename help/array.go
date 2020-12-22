package help

func ArrayUnique(ids []string) []string {
	m := make(map[string]bool)
	for _, id := range ids {
		if id == "" {
			continue
		}
		m[id] = true
	}
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}