package go_sensitive_word

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 敏感词大小写验证
func TestUppercaseLowercase(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)
	assert.NoError(t, err, "初始化敏感词过滤器失败")

	err = filter.Store.AddWord("sb")
	assert.NoError(t, err, "添加敏感词失败")

	sensitiveText := "测试Sb和sB还有SB还有sb"

	// 是否有敏感词
	res1 := filter.IsSensitive(sensitiveText)
	assert.True(t, res1, "应检测到敏感词")

	// 找到一个敏感词
	res2 := filter.FindOne(sensitiveText)
	assert.Equal(t, "sb", res2, "应检测出 sb 为敏感词")

	// 找到所有敏感词
	res3 := filter.FindAll(sensitiveText)
	assert.Contains(t, res3, "sb", "FindAll 应包含 sb")

	// 找到所有敏感词及出现次数
	res4 := filter.FindAllCount(sensitiveText)
	assert.GreaterOrEqual(t, res4["sb"], 1, "sb 至少应出现一次")

	// 和谐敏感词
	res5 := filter.Replace(sensitiveText, '*')
	assert.NotContains(t, res5, "sb", "Replace 后结果中不应包含 sb")

	// 过滤敏感词
	res6 := filter.Remove(sensitiveText)
	assert.NotContains(t, res6, "sb", "Remove 后结果中不应包含 sb")

	// 输出内容
	fmt.Printf("res1: %v \n", res1)
	fmt.Printf("res2: %v \n", res2)
	fmt.Printf("res3: %v \n", res3)
	fmt.Printf("res4: %v \n", res4)
	fmt.Printf("res5: %v \n", res5)
	fmt.Printf("res6: %v \n", res6)
}

// DFA算法敏感词检测
func TestDFA(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)
	assert.NoError(t, err, "初始化敏感词过滤器失败")

	err = filter.LoadDictEmbed(
		DictCovid19,
		DictOther,
		DictReactionary,
		DictViolence,
		DictPeopleLife,
		DictPornography,
		DictAdditional,
		DictCorruption,
		DictTemporaryTencent,
	)
	assert.NoError(t, err, "加载词库失败")

	err = filter.Store.AddWord("测试1", "测试2", "成小王")
	assert.NoError(t, err, "添加自定义敏感词失败")

	// 等待协程处理添加的敏感词
	time.Sleep(50 * time.Millisecond)

	sensitiveText := "成小王微笑着对毒品销售说，我认为台湾国的人有点意思"

	// 是否有敏感词
	res1 := filter.IsSensitive(sensitiveText)
	fmt.Printf("res1: %v \n", res1)
	assert.True(t, res1, "应检测到敏感词")

	// 找到一个敏感词
	res2 := filter.FindOne(sensitiveText)
	fmt.Printf("res2: %v \n", res2)
	assert.NotEmpty(t, res2, "应至少检测出一个敏感词")

	// 找到所有敏感词
	res3 := filter.FindAll(sensitiveText)
	fmt.Printf("res3: %v \n", res3)
	assert.NotEmpty(t, res3, "FindAll 应返回敏感词列表")
	assert.Contains(t, res3, "成小王", "应检测到 成小王")
	assert.Contains(t, res3, "毒品", "应检测到 毒品")

	// 找到所有敏感词及出现次数
	res4 := filter.FindAllCount(sensitiveText)
	fmt.Printf("res4: %v \n", res4)
	assert.GreaterOrEqual(t, res4["毒品"], 1, "毒品 应出现至少一次")

	// 和谐敏感词
	res5 := filter.Replace(sensitiveText, '*')
	fmt.Printf("res5: %v \n", res5)
	for _, w := range res3 {
		assert.NotContains(t, res5, w, fmt.Sprintf("Replace 后不应包含 %s", w))
	}

	// 过滤敏感词
	res6 := filter.Remove(sensitiveText)
	fmt.Printf("res6: %v \n", res6)
	for _, w := range res3 {
		assert.NotContains(t, res6, w, fmt.Sprintf("Remove 后不应包含 %s", w))
	}
}

// AC自动机算法敏感词检测
func TestAC(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterAc},
	)
	assert.NoError(t, err, "初始化敏感词过滤器失败")

	err = filter.LoadDictEmbed(
		DictCovid19,
		DictOther,
		DictReactionary,
		DictViolence,
		DictPeopleLife,
		DictPornography,
		DictAdditional,
		DictCorruption,
		DictTemporaryTencent,
	)
	assert.NoError(t, err, "加载词库失败")

	err = filter.Store.AddWord("测试1", "测试2", "成小王")
	assert.NoError(t, err, "添加自定义敏感词失败")

	// 等待协程处理添加的敏感词
	time.Sleep(50 * time.Millisecond)

	sensitiveText := "成小王微笑着对毒品销售说，我认为台湾国的人有点意思"

	// 是否有敏感词
	res1 := filter.IsSensitive(sensitiveText)
	assert.True(t, res1, "应检测到敏感词")

	// 找到一个敏感词
	res2 := filter.FindOne(sensitiveText)
	assert.NotEmpty(t, res2, "应至少检测出一个敏感词")

	// 找到所有敏感词
	res3 := filter.FindAll(sensitiveText)
	assert.NotEmpty(t, res3, "FindAll 应返回敏感词列表")
	assert.Contains(t, res3, "成小王")
	assert.Contains(t, res3, "毒品")

	// 找到所有敏感词及出现次数
	res4 := filter.FindAllCount(sensitiveText)
	assert.GreaterOrEqual(t, res4["毒品"], 1, "毒品 应出现至少一次")

	// 和谐敏感词
	res5 := filter.Replace(sensitiveText, '*')
	for _, w := range res3 {
		assert.NotContains(t, res5, w, fmt.Sprintf("Replace 后不应包含 %s", w))
	}

	// 过滤敏感词
	res6 := filter.Remove(sensitiveText)
	for _, w := range res3 {
		assert.NotContains(t, res6, w, fmt.Sprintf("Remove 后不应包含 %s", w))
	}

	// 输出内容
	fmt.Printf("res1: %v \n", res1)
	fmt.Printf("res2: %v \n", res2)
	fmt.Printf("res3: %v \n", res3)
	fmt.Printf("res4: %v \n", res4)
	fmt.Printf("res5: %v \n", res5)
	fmt.Printf("res6: %v \n", res6)
}

// HTTP远程加载词库测试
func TestLoadDictHttp(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)
	assert.NoError(t, err, "初始化敏感词过滤器失败")

	err = filter.LoadDictHttp("https://res.yongwang.lu/github/go-sensitive-word/demo.txt")
	if err != nil {
		log.Fatalf("加载HTTP词库错误, err:%v", err)
		return
	}

	// 等待协程处理添加的敏感词
	time.Sleep(50 * time.Millisecond)

	sensitiveText := "我们晚点要一起做Http测试"

	res3 := filter.FindAll(sensitiveText)
	assert.NotEmpty(t, res3, "FindAll 应返回敏感词列表")
	assert.Contains(t, res3, "http测试")

	// 输出内容
	fmt.Printf("res3: %v \n", res3)
}
