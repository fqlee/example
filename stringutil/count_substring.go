package stringutil

// CountSubstring 计算子字符串在原字符串中出现的次数
func CountSubstring(s, substr string) int {
    count := 0
    start := 0
    for {
        pos := strings.Index(s[start:], substr)
        if pos == -1 {
            break
        }
        count++
        start += pos + 1
    }
    return count
}    
