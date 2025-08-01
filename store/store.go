package store

import "io"

type (
	Store interface {
		// LoadDictPath 从指定的本地路径加载词库文件（可传多个路径）
		LoadDictPath(path ...string) error
		// LoadDictEmbed （推荐）加载嵌入式词库内容（如 go:embed 提供的字符串）
		LoadDictEmbed(contents ...string) error
		// LoadDictHttp 从远程 URL 加载词库内容（支持多个 URL）
		LoadDictHttp(url ...string) error
		// LoadDict 从 io.Reader 加载词库内容（按行读取）
		LoadDict(reader io.Reader) error
		// ReadChan 返回一个通道，逐个输出当前存储中的所有敏感词（可用于异步加载到过滤器）
		ReadChan() <-chan string
		// ReadString 以字符串数组形式返回当前所有敏感词
		ReadString() []string
		// GetAddChan 获取新增敏感词的事件通道（用于实时监听词库变更）
		GetAddChan() <-chan string
		// GetDelChan 获取删除敏感词的事件通道（用于实时监听词库变更）
		GetDelChan() <-chan string
		// AddWord 添加一个或多个敏感词
		AddWord(words ...string) error
		// DelWord 删除一个或多个敏感词
		DelWord(words ...string) error
	}
)

// 接口实现验证
var _ Store = (*MemoryModel)(nil)
