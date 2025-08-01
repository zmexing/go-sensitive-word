package go_sensitive_word

import _ "embed"

// StoreMemory 类型常量定义
// 当前仅支持内存存储（StoreMemory），后续可扩展为 Redis、文件存储等。
const (
	StoreMemory = iota // 内存模式词库（默认）
)

// FilterDfa 类型常量定义
// 当前仅支持 DFA 算法（FilterDfa），后续可支持 Trie、AC自动机、正则等。
const (
	FilterDfa = iota // DFA 敏感词过滤算法（默认）
)

// StoreOption 定义了词库存储的配置选项
// Type 字段用于指定词库的存储实现方式，如内存、Redis、文件等。
type StoreOption struct {
	Type uint32 // 存储类型标识，例如 StoreMemory
}

// FilterOption 定义了敏感词过滤器的配置选项
// Type 字段用于指定过滤算法的实现方式，如 DFA、Trie、正则等。
type FilterOption struct {
	Type uint32 // 过滤器类型标识，例如 FilterDfa
}

// 内置敏感词词库（通过 go:embed 嵌入编译时）
// 这些变量可直接用于调用 LoadDictEmbed 加载内置词库内容，无需读取本地文件。
// 可按需选择加载不同类别的敏感词，例如政治类、暴恐类、色情类、贪腐类等。
//
// 使用示例：
//
//	err := manager.LoadDictEmbed(DictCovid19, DictPornography)
var (
	//go:embed text/COVID-19词库.txt
	DictCovid19 string
	//go:embed text/其他词库.txt
	DictOther string
	//go:embed text/反动词库.txt
	DictReactionary string
	//go:embed text/暴恐词库.txt
	DictViolence string
	//go:embed text/民生词库.txt
	DictPeopleLife string
	//go:embed text/色情词库.txt
	DictPornography string
	//go:embed text/补充词库.txt
	DictAdditional string
	//go:embed text/贪腐词库.txt
	DictCorruption string
	//go:embed text/零时-Tencent.txt
	DictTemporaryTencent string
)
