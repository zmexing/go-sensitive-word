package go_sensitive_word

import "regexp"

// HasEmail 判断字符串中是否存在邮箱地址
func HasEmail(s string) bool {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	return emailRegex.MatchString(s)
}

// MaskEmail 将字符串中存在的邮箱地址替换成 "*"
func MaskEmail(s string) string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	masked := emailRegex.ReplaceAllStringFunc(s, func(email string) string {
		return "***"
	})
	return masked
}

// HasURL 判断字符串中是否存在网址
func HasURL(s string) bool {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	return urlRegex.MatchString(s)
}

// MaskURL 将字符串中存在的网址替换成 "*"
func MaskURL(s string) string {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	masked := urlRegex.ReplaceAllStringFunc(s, func(url string) string {
		return "***"
	})
	return masked
}

// HasDigit 判断字符串中是否存在指定个数的数字（大于等于该数字）
func HasDigit(s string, count int) bool {
	digitRegex := regexp.MustCompile(`\d`)
	matches := digitRegex.FindAllString(s, -1)
	return len(matches) >= count
}

// MaskDigit 将字符串中存在的数字替换成 "*"
func MaskDigit(s string) string {
	digitRegex := regexp.MustCompile(`\d`)
	masked := digitRegex.ReplaceAllStringFunc(s, func(digit string) string {
		return "*"
	})
	return masked
}

// HasWechatID 判断字符串中是否存在微信号
func HasWechatID(s string) bool {
	wechatRegex := regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_-]{5,19}`)
	return wechatRegex.MatchString(s)
}

// MaskWechatID 将字符串中存在的微信号替换成 "*"
func MaskWechatID(s string) string {
	wechatRegex := regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_-]{5,19}`)
	masked := wechatRegex.ReplaceAllStringFunc(s, func(wechatID string) string {
		return "***"
	})
	return masked
}
