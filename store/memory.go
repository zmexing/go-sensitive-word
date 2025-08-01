package store

import (
	"bufio"
	"errors"
	"github.com/imroc/req/v3"
	cmap "github.com/orcaman/concurrent-map/v2"
	"io"
	"net/http"
	"os"
	"strings"
)

// MemoryModel 使用并发 map 实现的内存词库
type MemoryModel struct {
	store   cmap.ConcurrentMap[string, struct{}]
	addChan chan string
	delChan chan string
}

// NewMemoryModel 创建新的内存模型
func NewMemoryModel() *MemoryModel {
	return &MemoryModel{
		store:   cmap.New[struct{}](),
		addChan: make(chan string),
		delChan: make(chan string),
	}
}

// 从本地路径加载词库文件
func (m *MemoryModel) LoadDictPath(paths ...string) error {
	for _, path := range paths {
		err := func(path string) error {
			f, err := os.Open(path)
			defer func(f *os.File) {
				_ = f.Close()
			}(f)
			if err != nil {
				return err
			}

			return m.LoadDict(f)
		}(path)
		if err != nil {
			return err
		}
	}

	return nil
}

// 加载嵌入式文本词库（go:embed）
func (m *MemoryModel) LoadDictEmbed(contents ...string) error {
	for _, con := range contents {
		reader := strings.NewReader(con)
		if err := m.LoadDict(reader); err != nil {
			return err
		}
	}

	return nil
}

// 从远程 HTTP 地址加载词库
func (m *MemoryModel) LoadDictHttp(urls ...string) error {
	for _, url := range urls {
		err := func(url string) error {
			httpRes, err := req.Get(url)
			if err != nil {
				return err
			}
			if httpRes == nil {
				return errors.New("nil http response")
			}
			if httpRes.StatusCode != http.StatusOK {
				return errors.New(httpRes.GetStatus())
			}

			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(httpRes.Body)

			return m.LoadDict(httpRes.Body)
		}(url)
		if err != nil {
			return err
		}
	}

	return nil
}

// 读取词库（按行解析）
func (m *MemoryModel) LoadDict(reader io.Reader) error {
	buf := bufio.NewReader(reader)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		m.store.Set(string(line), struct{}{})
		m.addChan <- string(line)
	}

	return nil
}

// 返回所有敏感词的读取通道（可用于初始化加载）
func (m *MemoryModel) ReadChan() <-chan string {
	ch := make(chan string)

	go func() {
		for key := range m.store.Items() {
			ch <- key
		}
		close(ch)
	}()

	return ch
}

// 获取所有敏感词（字符串数组）
func (m *MemoryModel) ReadString() []string {
	res := make([]string, 0, m.store.Count())

	for key := range m.store.Items() {
		res = append(res, key)
	}

	return res
}

// 获取新增词通道
func (m *MemoryModel) GetAddChan() <-chan string {
	return m.addChan
}

// 获取删除词通道
func (m *MemoryModel) GetDelChan() <-chan string {
	return m.delChan
}

// 添加自定义敏感词
func (m *MemoryModel) AddWord(words ...string) error {
	for _, word := range words {
		m.store.Set(word, struct{}{})
		m.addChan <- word
	}

	return nil
}

// 删除敏感词（敏感词加白名单）
func (m *MemoryModel) DelWord(words ...string) error {
	for _, word := range words {
		m.store.Remove(word)
		m.delChan <- word
	}

	return nil
}
