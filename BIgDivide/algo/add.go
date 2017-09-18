package algo

import "fmt"

func Add(a, b string) string {
	lenA := len(a)
	lenB := len(b)
	var lenth int
	var start int
	var result []byte
	var toAdd string
	var add string
	if lenA > lenB {
		start = lenA - lenB
		lenth = lenA
		result = make([]byte, lenA+1)
		add = a
		toAdd = b
	} else {
		start = lenB - lenA
		lenth = lenB
		result = make([]byte, lenB+1)
		add = b
		toAdd = a
	}

	for i := 0; i < start; i++ {
		result[i+1] = add[i]
	}
	for i := start; i < lenth; i++ {
		result[i+1] = add[i] + toAdd[i-start] - '0'
	}
	fmt.Println(fmt.Sprintf("这里的result=%v",result))
	resultLen := len(result)
	for i := 0; i < resultLen-1; i++ {
		index := resultLen - i - 1
		v := result[index] - '0'

		if v > 9 {
			resultValue := v % 10
			resultJinwei := v / 10
			result[index] = resultValue
			result[index-1] += resultJinwei
		}
	}

	i := 0
	for result[i] == 0 {
		i++
	}
	final := result[i:]
	fmt.Println(fmt.Sprintf("大数相加final=%s,result=%s", final, result))
	return string(final)
}
