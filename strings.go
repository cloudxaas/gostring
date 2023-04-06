package cxstrings

func TrimStringFromQuery(s, substr string) string {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return s[:i]
		}
	}
	return s
}
