package go_sensitive_word

import (
	"fmt"
	"log"
	"testing"
)

func TestFilter1(t *testing.T) {
	filter := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)

	// 加载敏感词库
	err := filter.Store.LoadDictPath("./text/dict2.txt")
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	// 动态自定义敏感词
	err = filter.Store.AddWord("测试1", "测试2", "李先生")
	if err != nil {
		fmt.Println(err)
		return
	}

	sensitiveText := "李先生微笑着对毒品销售说，我认为台湾国的人有点意思"

	// 是否有敏感词
	res1 := filter.IsSensitive(sensitiveText)
	fmt.Printf("res1: %v \n", res1)

	// 找到一个敏感词
	res2 := filter.FindOne(sensitiveText)
	fmt.Printf("res2: %v \n", res2)

	// 找到所有敏感词
	res3 := filter.FindAll(sensitiveText)
	fmt.Printf("res3: %v \n", res3)

	// 找到所有敏感词及出现次数
	res4 := filter.FindAllCount(sensitiveText)
	fmt.Printf("res4: %v \n", res4)

	// 和谐敏感词
	res5 := filter.Replace(sensitiveText, '*')
	fmt.Printf("res5: %v \n", res5)

	// 过滤铭感词
	res6 := filter.Remove(sensitiveText)
	fmt.Printf("res6: %v \n", res6)
}

func BenchmarkIsSensitive(b *testing.B) {
	filter := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)

	err := filter.Store.LoadDictPath("./text/dict2.txt")
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	err = filter.Store.AddWord("测试1", "测试2")
	if err != nil {
		fmt.Println(err)
		return
	}

	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.IsSensitive(sensitiveText)
	}
}

func BenchmarkReplace(b *testing.B) {
	filter := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)

	err := filter.Store.LoadDictPath("./text/dict2.txt")
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	err = filter.Store.AddWord("测试1", "测试2")
	if err != nil {
		fmt.Println(err)
		return
	}

	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.Replace(sensitiveText, '*')
	}
}
