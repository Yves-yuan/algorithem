package algo

import (
	"bytes"
	"fmt"
)

//去除大数前面的0
func removeZero(a string) string {
	i := 0
	for '0' == a[i] && i < len(a)-1 {
		i++
	}
	a = a[i:]
	return a
}

//判断大小,false代表a小于b,true代表a大于等于b
func Compare(a, b string) bool {
	a = removeZero(a)
	b = removeZero(b)

	len1 := len(a)
	len2 := len(b)

	if len1 < len2 {
		return false
	} else if len1 == len2 && a < b {
		return false
	}
	return true
}

//大数相减a减b
func Subtract(a, b string) string {
	if a == b {
		return "0"
	}

	len1 := len(a)
	len2 := len(b)

	runeA := []rune(a)
	for i := len2 - 1; i >= 0; i-- {
		runeA[len1-1-i] = rune('0' + a[len1-1-i] - b[len2-1-i])
	}

	for i := len1 - 1; i >= 0; i-- {
		if runeA[i] < '0' {
			runeA[i] += 10
			runeA[i-1]--
		}
	}

	i := 0
	for runeA[i] == '0' {
		i++
	}

	runeA = runeA[i:]

	return string(runeA)
}

//大数相除
func BigDivide(a, b string) (quotient, remainder string) {
	var result bytes.Buffer
	var s bytes.Buffer

	var (
		count int
	)

	if "0" == b {
		fmt.Println("除数为0")
		return
	}

	if false == Compare(a,b) {
		quotient = "0"
		remainder = a
		return
	}

	len := len(a)
	for i := 0; i < len; i++ {
		count = 0
		s.WriteByte(a[i])
		str := s.String()
		for Compare(str,b){
			str = Subtract(str, b)
			count++
		}
		s.Reset()
		s.WriteString(str)
		result.WriteByte(byte('0' + count))
	}

	resultStr := result.String()
	resultStr = removeZero(resultStr)

	sStr := s.String()
	sStr = removeZero(sStr)

	return resultStr, sStr
}
