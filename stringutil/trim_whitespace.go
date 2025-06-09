package stringutil

import "strings"

// TrimWhitespace 去除字符串中的所有空白字符
func TrimWhitespace(s string) string {
    return strings.Join(strings.Fields(s), "")
}    
