package store

import "io"

type (
	Store interface {
		LoadDictPath(path ...string) error
		LoadDictHttp(url ...string) error
		LoadDict(reader io.Reader) error
		ReadChan() <-chan string
		ReadString() []string
		GetAddChan() <-chan string
		GetDelChan() <-chan string
		AddWord(words ...string) error
		DelWord(words ...string) error
	}
)

// 接口实现验证
var _ Store = (*MemoryModel)(nil)
