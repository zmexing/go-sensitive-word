package filter

// DFA 树节点结构
type dfaNode struct {
	children map[rune]*dfaNode // 子节点
	isLeaf   bool              // 是否为词尾
}

// DfaModel 是基于 DFA 的敏感词匹配器
func newDfaNode() *dfaNode {
	return &dfaNode{
		children: make(map[rune]*dfaNode),
		isLeaf:   false,
	}
}

type DfaModel struct {
	root *dfaNode
}

func NewDfaModel() *DfaModel {
	return &DfaModel{
		root: newDfaNode(),
	}
}

// 添加多个词
func (m *DfaModel) AddWords(words ...string) {
	for _, word := range words {
		m.AddWord(word)
	}
}

// 添加单个词到 DFA 树中
func (m *DfaModel) AddWord(word string) {
	now := m.root
	runes := []rune(word)

	for _, r := range runes {
		if next, ok := now.children[r]; ok {
			now = next
		} else {
			next = newDfaNode()
			now.children[r] = next
			now = next
		}
	}

	now.isLeaf = true
}

// 删除多个词
func (m *DfaModel) DelWords(words ...string) {
	for _, word := range words {
		m.DelWord(word)
	}
}

// 删除单个词（仅支持叶子节点剪枝）
func (m *DfaModel) DelWord(word string) {
	var lastLeaf *dfaNode
	var lastLeafNextRune rune
	now := m.root
	runes := []rune(word)

	for _, r := range runes {
		if next, ok := now.children[r]; !ok {
			return
		} else {
			if now.isLeaf {
				lastLeaf = now
				lastLeafNextRune = r
			}
			now = next
		}
	}

	delete(lastLeaf.children, lastLeafNextRune)
}

// 监听新增和删除通道
func (m *DfaModel) Listen(addChan, delChan <-chan string) {
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

// 查找文本中所有敏感词
func (m *DfaModel) FindAll(text string) []string {
	var matches []string // stores words that match in dict
	var found bool       // if current rune in node's map
	var now *dfaNode     // current node

	start := 0
	parent := m.root
	runes := []rune(text)
	length := len(runes)

	for pos := 0; pos < length; pos++ {
		now, found = parent.children[runes[pos]]

		if !found {
			parent = m.root
			pos = start
			start++
			continue
		}

		if now.isLeaf && start <= pos {
			matches = append(matches, string(runes[start:pos+1]))
		}

		if pos == length-1 {
			parent = m.root
			pos = start
			start++
			continue
		}

		parent = now
	}

	var res []string
	set := make(map[string]struct{})

	for _, word := range matches {
		if _, ok := set[word]; !ok {
			set[word] = struct{}{}
			res = append(res, word)
		}
	}

	return res
}

// 查找所有敏感词及其出现次数
func (m *DfaModel) FindAllCount(text string) map[string]int {
	res := make(map[string]int)
	var found bool
	var now *dfaNode

	start := 0
	parent := m.root
	runes := []rune(text)
	length := len(runes)

	for pos := 0; pos < length; pos++ {
		now, found = parent.children[runes[pos]]

		if !found {
			parent = m.root
			pos = start
			start++
			continue
		}

		if now.isLeaf && start <= pos {
			res[string(runes[start:pos+1])]++
		}

		if pos == length-1 {
			parent = m.root
			pos = start
			start++
			continue
		}

		parent = now
	}

	return res
}

// 查找一个敏感词（命中第一个即返回）
func (m *DfaModel) FindOne(text string) string {
	var found bool
	var now *dfaNode

	start := 0
	parent := m.root
	runes := []rune(text)
	length := len(runes)

	for pos := 0; pos < length; pos++ {
		now, found = parent.children[runes[pos]]

		if !found || (!now.isLeaf && pos == length-1) {
			parent = m.root
			pos = start
			start++
			continue
		}

		if now.isLeaf && start <= pos {
			return string(runes[start : pos+1])
		}

		parent = now
	}

	return ""
}

// 判断文本中是否包含敏感词
func (m *DfaModel) IsSensitive(text string) bool {
	return m.FindOne(text) != ""
}

// 将敏感词替换为指定字符（如 *）
func (m *DfaModel) Replace(text string, repl rune) string {
	var found bool
	var now *dfaNode

	start := 0
	parent := m.root
	runes := []rune(text)
	length := len(runes)

	for pos := 0; pos < length; pos++ {
		now, found = parent.children[runes[pos]]

		if !found || (!now.isLeaf && pos == length-1) {
			parent = m.root
			pos = start
			start++
			continue
		}

		if now.isLeaf && start <= pos {
			for i := start; i <= pos; i++ {
				runes[i] = repl
			}
		}

		parent = now
	}

	return string(runes)
}

// 将敏感词从文本中完全移除
func (m *DfaModel) Remove(text string) string {
	var found bool
	var now *dfaNode

	start := 0 // 从文本的第几个文字开始匹配
	parent := m.root
	runes := []rune(text)
	length := len(runes)
	filtered := make([]rune, 0, length)

	for pos := 0; pos < length; pos++ {
		now, found = parent.children[runes[pos]]

		if !found || (!now.isLeaf && pos == length-1) {
			filtered = append(filtered, runes[start])
			parent = m.root
			pos = start
			start++
			continue
		}

		if now.isLeaf {
			start = pos + 1
			parent = m.root
		} else {
			parent = now
		}
	}

	filtered = append(filtered, runes[start:]...)

	return string(filtered)
}

// Listen 监听新增和删除通道（线程安全建议加锁）
func (m *AcModel) Listen(addChan, delChan <-chan string) {
	go func() {
		for word := range addChan {
			m.AddWord(word)
			m.Build() // 重建 fail 指针
		}
	}()
	go func() {
		for word := range delChan {
			m.DelWord(word)
			m.Build() // 重建 fail 指针
		}
	}()
}

// DelWord 删除一个敏感词
func (m *AcModel) DelWord(word string) {
	node := m.root
	path := []*acNode{}
	runes := []rune(word)

	for _, r := range runes {
		next, ok := node.children[r]
		if !ok {
			return
		}
		path = append(path, next)
		node = next
	}

	// 从 output 中删除这个词
	out := node.output
	newOut := out[:0]
	for _, w := range out {
		if w != word {
			newOut = append(newOut, w)
		}
	}
	node.output = newOut
	// 简化处理，不回溯剪枝节点，保留空节点
}
