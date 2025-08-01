package go_sensitive_word

import (
	"errors"
	"github.com/zmexing/go-sensitive-word/filter"
	"github.com/zmexing/go-sensitive-word/store"
)

// Manager 是敏感词过滤系统的核心结构，整合了词库存储和过滤算法
type Manager struct {
	store.Store   //  // 词库存储接口（支持内存、本地文件、远程等）
	filter.Filter // // 敏感词匹配算法接口（如 DFA）
}

// NewFilter 初始化过滤器和词库存储
// 参数：storeOption 指定存储方式，filterOption 指定过滤算法
func NewFilter(storeOption StoreOption, filterOption FilterOption) (*Manager, error) {
	var filterStore store.Store
	var myFilter filter.Filter

	switch storeOption.Type {
	case StoreMemory: // 使用内存词库
		filterStore = store.NewMemoryModel()
	default:
		return nil, errors.New("invalid store type")
	}

	switch filterOption.Type {
	case FilterDfa: // 使用 DFA 算法
		dfaModel := filter.NewDfaModel()
		// 启动监听协程，实时接收新增/删除词的通知
		go dfaModel.Listen(filterStore.GetAddChan(), filterStore.GetDelChan())
		myFilter = dfaModel
	default:
		return nil, errors.New("invalid filter type")
	}

	return &Manager{
		Store:  filterStore,
		Filter: myFilter,
	}, nil
}
