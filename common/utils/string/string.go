package string

import (
	"math/rand"
	"time"
)

func RemoveDuplicate(slc []string) []string {
	m := make(map[string]bool)
	for _, v := range slc {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

// 生成指定长度的随机字符串
func RandomStringGenerator(length int) string {
	rand.Seed(time.Now().UnixNano())

	// 定义可用字符集合
	charset := "0123456789"

	// 创建一个字符切片，用于存储随机生成的字符
	result := make([]byte, length)

	// 生成随机字符串
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result[i] = charset[randomIndex]
	}

	return string(result)
}
