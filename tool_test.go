package go_sensitive_word

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	input := "我的邮箱是example@example.com，你的是test@test.com。"

	if HasEmail(input) {
		masked := MaskEmail(input)
		fmt.Println("替换后的字符串:", masked)
	} else {
		fmt.Println("字符串中不存在邮箱。")
	}
}

func TestURL(t *testing.T) {
	input := "我的网址是http://example.com，你的是https://test.com。"

	if HasURL(input) {
		masked := MaskURL(input)
		fmt.Println("替换后的字符串:", masked)
	} else {
		fmt.Println("字符串中不存在网址。")
	}
}

func TestDigit(t *testing.T) {
	input := "我的手机号码是1234567890，你的是9876543210。"
	if HasDigit(input, 3) {
		masked := MaskDigit(input)
		fmt.Println("替换后的字符串:", masked)
	} else {
		fmt.Println("字符串中不存在指定个数的数字。")
	}
}

func TestWechat(t *testing.T) {
	input := "我的是my_wechat123，你的微信是your-wechat-789。"
	if HasWechatID(input) {
		masked := MaskWechatID(input)
		fmt.Println("替换后的字符串:", masked)
	} else {
		fmt.Println("字符串中不存在微信号。")
	}
}
