package help

// 去除slice中重复元素和空字符串
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

// SliceChunk 切分slice为相同size的slice
func ArrayChunk(data []string, chunkSize int) [][]string {
	var result [][]string
	if chunkSize <= 0 {
		return result
	}

	data = ArrayUnique(data)

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		if end > len(data) {
			end = len(data)
		}

		result = append(result, data[i:end])
	}

	return result
}