package store

import (
	"bufio"
	"errors"
	"github.com/imroc/req/v3"
	cmap "github.com/orcaman/concurrent-map/v2"
	"io"
	"net/http"
	"os"
)

type MemoryModel struct {
	store   cmap.ConcurrentMap[string, struct{}]
	addChan chan string
	delChan chan string
}

func NewMemoryModel() *MemoryModel {
	return &MemoryModel{
		store:   cmap.New[struct{}](),
		addChan: make(chan string),
		delChan: make(chan string),
	}
}

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

func (m *MemoryModel) ReadString() []string {
	res := make([]string, 0, m.store.Count())

	for key := range m.store.Items() {
		res = append(res, key)
	}

	return res
}

func (m *MemoryModel) GetAddChan() <-chan string {
	return m.addChan
}

func (m *MemoryModel) GetDelChan() <-chan string {
	return m.delChan
}

func (m *MemoryModel) AddWord(words ...string) error {
	for _, word := range words {
		m.store.Set(word, struct{}{})
		m.addChan <- word
	}

	return nil
}

func (m *MemoryModel) DelWord(words ...string) error {
	for _, word := range words {
		m.store.Remove(word)
		m.delChan <- word
	}

	return nil
}
