package filter

import "container/list"

// acNode 是 Aho-Corasick 的节点
type acNode struct {
	children map[rune]*acNode
	fail     *acNode
	output   []string
}

// AcModel 是基于 Aho-Corasick 的敏感词匹配器
type AcModel struct {
	root *acNode
}

// NewAcModel 创建 Aho-Corasick 模型
func NewAcModel() *AcModel {
	return &AcModel{
		root: &acNode{
			children: make(map[rune]*acNode),
			fail:     nil,
			output:   nil,
		},
	}
}

// AddWords 批量添加敏感词
func (m *AcModel) AddWords(words ...string) {
	for _, word := range words {
		m.AddWord(word)
	}
}

// AddWord 添加一个敏感词
func (m *AcModel) AddWord(word string) {
	node := m.root
	for _, r := range []rune(word) {
		if next, ok := node.children[r]; ok {
			node = next
		} else {
			newNode := &acNode{
				children: make(map[rune]*acNode),
			}
			node.children[r] = newNode
			node = newNode
		}
	}
	node.output = append(node.output, word)
}

// Build 构建 fail 指针
func (m *AcModel) Build() {
	queue := list.New()
	// 初始化第一层
	for _, child := range m.root.children {
		child.fail = m.root
		queue.PushBack(child)
	}

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(*acNode)
		for r, next := range front.children {
			queue.PushBack(next)

			// 回溯 fail 指针
			failNode := front.fail
			for failNode != nil {
				if child, ok := failNode.children[r]; ok {
					next.fail = child
					next.output = append(next.output, child.output...)
					break
				}
				failNode = failNode.fail
			}
			if failNode == nil {
				next.fail = m.root
			}
		}
	}
}

// FindAll 查找所有敏感词
func (m *AcModel) FindAll(text string) []string {
	node := m.root
	var result []string
	for _, r := range []rune(text) {
		for node != m.root && node.children[r] == nil {
			node = node.fail
		}
		if next, ok := node.children[r]; ok {
			node = next
		}
		if len(node.output) > 0 {
			result = append(result, node.output...)
		}
	}
	// 去重
	uniq := make(map[string]struct{})
	var res []string
	for _, w := range result {
		if _, ok := uniq[w]; !ok {
			uniq[w] = struct{}{}
			res = append(res, w)
		}
	}
	return res
}

// FindAllCount 查找所有敏感词及次数
func (m *AcModel) FindAllCount(text string) map[string]int {
	node := m.root
	res := make(map[string]int)
	for _, r := range []rune(text) {
		for node != m.root && node.children[r] == nil {
			node = node.fail
		}
		if next, ok := node.children[r]; ok {
			node = next
		}
		for _, out := range node.output {
			res[out]++
		}
	}
	return res
}

// FindOne 找到第一个敏感词
func (m *AcModel) FindOne(text string) string {
	node := m.root
	for _, r := range []rune(text) {
		for node != m.root && node.children[r] == nil {
			node = node.fail
		}
		if next, ok := node.children[r]; ok {
			node = next
		}
		if len(node.output) > 0 {
			return node.output[0]
		}
	}
	return ""
}

// IsSensitive 是否包含敏感词
func (m *AcModel) IsSensitive(text string) bool {
	return m.FindOne(text) != ""
}

// Replace 替换敏感词
func (m *AcModel) Replace(text string, repl rune) string {
	runes := []rune(text)
	node := m.root
	for i, r := range runes {
		for node != m.root && node.children[r] == nil {
			node = node.fail
		}
		if next, ok := node.children[r]; ok {
			node = next
		}
		if len(node.output) > 0 {
			for _, w := range node.output {
				start := i - len([]rune(w)) + 1
				for j := start; j <= i; j++ {
					runes[j] = repl
				}
			}
		}
	}
	return string(runes)
}

// Remove 移除敏感词
func (m *AcModel) Remove(text string) string {
	runes := []rune(text)
	node := m.root
	mask := make([]bool, len(runes))
	for i, r := range runes {
		for node != m.root && node.children[r] == nil {
			node = node.fail
		}
		if next, ok := node.children[r]; ok {
			node = next
		}
		if len(node.output) > 0 {
			for _, w := range node.output {
				start := i - len([]rune(w)) + 1
				for j := start; j <= i; j++ {
					mask[j] = true
				}
			}
		}
	}
	var filtered []rune
	for i, r := range runes {
		if !mask[i] {
			filtered = append(filtered, r)
		}
	}
	return string(filtered)
}
