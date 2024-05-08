package filter

type (
	Filter interface {
		// FindAll 找到所有敏感词
		FindAll(text string) []string
		// FindAllCount 找到所有敏感词及出现次数
		FindAllCount(text string) map[string]int
		// FindOne 找到一个敏感词
		FindOne(text string) string
		// IsSensitive 是否有敏感词
		IsSensitive(text string) bool
		// Replace 和谐敏感词
		Replace(text string, repl rune) string
		// Remove 过滤铭感词
		Remove(text string) string
	}
)

// 接口实现验证
var _ Filter = (*DfaModel)(nil)
