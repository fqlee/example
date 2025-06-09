// Package stringutil 提供了常用的字符串处理函数
package stringutil

// Reverse 将字符串反转后返回
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}    
