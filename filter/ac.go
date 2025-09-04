package filter

import "strings"

// AC 自动机节点结构
type acNode struct {
	children map[rune]*acNode // 子节点
	fail     *acNode          // 失败指针
	output   []string         // 该节点对应的完整单词列表
}

// 创建新的 AC 节点
func newAcNode() *acNode {
	return &acNode{
		children: make(map[rune]*acNode),
		fail:     nil,
		output:   make([]string, 0),
	}
}

// AcModel 是基于 Aho-Corasick 的敏感词匹配器
type AcModel struct {
	root  *acNode
	built bool // 标记是否已构建失败指针
}

// NewAcModel 创建新的 AC 自动机
func NewAcModel() *AcModel {
	return &AcModel{
		root:  newAcNode(),
		built: false,
	}
}

// AddWord 添加单个词到 AC 自动机中
func (m *AcModel) AddWord(word string) {
	if word == "" {
		return
	}

	m.built = false // 重置构建状态
	now := m.root
	for _, r := range []rune(word) {
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			next = newAcNode()
			now.children[r] = next
			now = next
		}
	}
	// 直接追加到 output，表示此节点有对应敏感词
	now.output = append(now.output, word)
}

// AddWords 添加多个词
func (m *AcModel) AddWords(output ...string) {
	for _, word := range output {
		m.AddWord(word)
	}
}

// DelWord 删除单个词
func (m *AcModel) DelWord(word string) {
	if word == "" {
		return
	}

	m.built = false // 重置构建状态
	now := m.root
	for _, r := range []rune(word) {
		next, ok := now.children[r]
		if !ok {
			return // 词不存在
		}
		now = next
	}
	// 从 output 列表中移除
	for i, w := range now.output {
		if w == word {
			now.output = append(now.output[:i], now.output[i+1:]...)
			break
		}
	}
}

// Deloutput 删除多个词
func (m *AcModel) Deloutput(output ...string) {
	for _, word := range output {
		m.DelWord(word)
	}
}

// buildFailurePointer 构建失败指针
func (m *AcModel) buildFailurePointer() {
	if m.built {
		return
	}

	queue := make([]*acNode, 0)
	// 第一层节点的失败指针指向根节点
	for _, child := range m.root.children {
		child.fail = m.root
		queue = append(queue, child)
	}

	// BFS 构建失败指针
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for char, child := range current.children {
			queue = append(queue, child)

			temp := current.fail
			for temp != nil && temp.children[char] == nil {
				temp = temp.fail
			}

			if temp == nil {
				child.fail = m.root
			} else {
				child.fail = temp.children[char]
				// 继承 fail 节点的输出
				child.output = append(child.output, child.fail.output...)
			}
		}
	}

	m.built = true
}

// Listen 监听新增和删除通道
func (m *AcModel) Listen(addChan, delChan <-chan string) {
	go func() {
		for word := range addChan {
			m.AddWord(word)
		}
	}()
	go func() {
		for word := range delChan {
			m.DelWord(word)
		}
	}()
}

// FindAll 查找文本中所有敏感词
func (m *AcModel) FindAll(text string) []string {
	m.buildFailurePointer()

	var matches []string
	seen := make(map[string]struct{})
	now := m.root
	for _, r := range []rune(text) {
		for now != m.root && now.children[r] == nil {
			now = now.fail
		}
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			now = m.root
		}
		for _, w := range now.output {
			if _, ok := seen[w]; !ok {
				seen[w] = struct{}{}
				matches = append(matches, w)
			}
		}
	}
	return matches
}

// FindAllCount 查找所有敏感词及其出现次数
func (m *AcModel) FindAllCount(text string) map[string]int {
	m.buildFailurePointer()

	counts := make(map[string]int)
	now := m.root
	for _, r := range []rune(text) {
		for now != m.root && now.children[r] == nil {
			now = now.fail
		}
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			now = m.root
		}
		for _, w := range now.output {
			counts[w]++
		}
	}
	return counts
}

// FindOne 查找一个敏感词（优先返回最长匹配）
func (m *AcModel) FindOne(text string) string {
	m.buildFailurePointer()

	now := m.root
	for _, r := range []rune(text) {
		for now != m.root && now.children[r] == nil {
			now = now.fail
		}
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			now = m.root
		}
		if len(now.output) > 0 {
			// 返回当前节点的最长匹配
			longest := now.output[0]
			for i := 1; i < len(now.output); i++ {
				if len([]rune(now.output[i])) > len([]rune(longest)) {
					longest = now.output[i]
				}
			}
			return longest
		}
	}
	return ""
}

// IsSensitive 判断文本中是否包含敏感词
func (m *AcModel) IsSensitive(text string) bool {
	return m.FindOne(text) != ""
}

// Replace 将敏感词替换为指定字符
func (m *AcModel) Replace(text string, repl rune) string {
	m.buildFailurePointer()

	runes := []rune(text)
	now := m.root
	for i, r := range runes {
		for now != m.root && now.children[r] == nil {
			now = now.fail
		}
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			now = m.root
		}
		for _, w := range now.output {
			wordLen := len([]rune(w))
			start := i - wordLen + 1
			if start >= 0 {
				for j := start; j <= i; j++ {
					runes[j] = repl
				}
			}
		}
	}
	return string(runes)
}

// Remove 将敏感词从文本中完全移除
func (m *AcModel) Remove(text string) string {
	matches := m.FindAll(text)
	// 按长度降序，先移除长词避免冲突
	for i := 0; i < len(matches); i++ {
		for j := i + 1; j < len(matches); j++ {
			if len([]rune(matches[i])) < len([]rune(matches[j])) {
				matches[i], matches[j] = matches[j], matches[i]
			}
		}
	}
	result := text
	for _, w := range matches {
		result = strings.ReplaceAll(result, w, "")
	}
	return result
}
