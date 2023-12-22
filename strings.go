package cxstrings

func TrimStringFromQuery(s, substr string) string {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return s[:i]
		}
	}
	return s
}

func Strpos(haystack, needle string, offset int) int {
    length := len(haystack)
    if length == 0 || offset >= length || -offset > length {
        return -1
    }
 
    if offset < 0 {
        offset += length
    }
 
    for i := offset; i <= length-len(needle); i++ {
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
 
    return -1
}
 
 
func Substr(str string, start uint, length int) string {
    if int(start) >= len(str) || length == 0 {
        return ""
    }
    if length < 0 || int(start)+length > len(str) {
        return str[start:]
    }
    return str[start : start+uint(length)]
}
