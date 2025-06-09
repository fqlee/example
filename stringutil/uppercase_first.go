package stringutil

import "unicode"

// UppercaseFirst 将字符串的首字母转换为大写
func UppercaseFirst(s string) string {
    if s == "" {
        return s
    }
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}    
