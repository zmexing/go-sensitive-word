package fetch

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// HttpConfig HTTP 请求配置
type HttpConfig struct {
	Method  string            // HTTP 方法 (GET, POST, PUT, DELETE 等)
	Headers map[string]string // 自定义请求头
	Body    io.Reader         // 请求体 (POST/PUT 等请求使用)
}

// DefaultHttpConfig 返回默认的 HTTP 配置
func DefaultHttpConfig() *HttpConfig {
	return &HttpConfig{
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Go Http Client By Sensitive Word",
		},
		Body: nil,
	}
}

// PostConfig 创建 POST 请求配置
func PostConfig(contentType string, body io.Reader) *HttpConfig {
	config := DefaultHttpConfig()
	config.Method = "POST"
	config.Body = body
	if contentType != "" {
		config.Headers["Content-Type"] = contentType
	}
	return config
}

// PostJsonConfig 创建 JSON POST 请求配置
func PostJsonConfig(jsonData []byte) *HttpConfig {
	config := DefaultHttpConfig()
	config.Method = "POST"
	config.Body = bytes.NewReader(jsonData)
	config.Headers["Content-Type"] = "application/json"
	return config
}

// PostFormConfig 创建表单 POST 请求配置
func PostFormConfig(formData []byte) *HttpConfig {
	config := DefaultHttpConfig()
	config.Method = "POST"
	config.Body = bytes.NewReader(formData)
	config.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	return config
}

// FetchUrl 发送 HTTP 请求并返回 HTTP 响应
func FetchUrl(url string, config *HttpConfig) (*http.Response, error) {
	if config == nil {
		config = DefaultHttpConfig()
	}

	// 创建请求
	req, err := http.NewRequest(config.Method, url, config.Body)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 添加自定义 Header
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	return resp, nil
}

// FetchUrlSimple 发送简单的 GET 请求（向后兼容）
func FetchUrlSimple(url string) (*http.Response, error) {
	return FetchUrl(url, DefaultHttpConfig())
}
