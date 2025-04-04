package ch01

import "testing"

func TestIotaConstSpec(t *testing.T) {
	// tips: 容易弄混地，itoa 是 C语言传统命名的整数转字符串的函数，即 integer to ASCII
	const (
		_ = iota // 常量自动递增声明
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)

	t.Log(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}
