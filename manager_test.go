package go_sensitive_word

import (
	"fmt"
	"log"
	"testing"
)

// 压力测试
// func BenchmarkIsSensitive(b *testing.B) {
// 	filter, err := NewFilter(
// 		StoreOption{Type: StoreMemory},
// 		FilterOption{Type: FilterDfa},
// 	)
//
// 	if err != nil {
// 		log.Fatalf("敏感词服务启动失败, err:%v", err)
// 		return
// 	}
//
// 	// 加载敏感词库
// 	err = filter.LoadDictEmbed(
// 		DictCovid19,
// 		DictOther,
// 		DictReactionary,
// 		DictViolence,
// 		DictPeopleLife,
// 		DictPornography,
// 		DictAdditional,
// 		DictCorruption,
// 		DictTemporaryTencent,
// 	)
// 	if err != nil {
// 		log.Fatalf("加载词库发生了错误, err:%v", err)
// 		return
// 	}
//
// 	err = filter.Store.AddWord("测试1", "测试2")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = filter.IsSensitive(sensitiveText)
// 	}
// }

// 压力测试
// func BenchmarkReplace(b *testing.B) {
// 	filter, err := NewFilter(
// 		StoreOption{Type: StoreMemory},
// 		FilterOption{Type: FilterDfa},
// 	)
//
// 	if err != nil {
// 		log.Fatalf("敏感词服务启动失败, err:%v", err)
// 		return
// 	}
//
// 	// 加载敏感词库
// 	err = filter.LoadDictEmbed(
// 		DictCovid19,
// 		DictOther,
// 		DictReactionary,
// 		DictViolence,
// 		DictPeopleLife,
// 		DictPornography,
// 		DictAdditional,
// 		DictCorruption,
// 		DictTemporaryTencent,
// 	)
// 	if err != nil {
// 		log.Fatalf("加载词库发生了错误, err:%v", err)
// 		return
// 	}
//
// 	err = filter.Store.AddWord("测试1", "测试2")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = filter.Replace(sensitiveText, '*')
// 	}
// }

// 敏感词检测
func TestDFA(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterDfa},
	)
	if err != nil {
		log.Fatalf("敏感词服务启动失败, err:%v", err)
		return
	}

	// 加载敏感词库
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
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	// 动态自定义敏感词
	err = filter.Store.AddWord("测试1", "测试2", "成小王")
	if err != nil {
		fmt.Println(err)
		return
	}

	sensitiveText := "成小王微笑着对毒品销售说，我认为台湾国的人有点意思"

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

func TestAC(t *testing.T) {
	filter, err := NewFilter(
		StoreOption{Type: StoreMemory},
		FilterOption{Type: FilterAc},
	)
	if err != nil {
		log.Fatalf("敏感词服务启动失败, err:%v", err)
		return
	}

	// 加载敏感词库
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
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	// 动态自定义敏感词
	err = filter.Store.AddWord("测试1", "测试2", "成小王")
	if err != nil {
		fmt.Println(err)
		return
	}

	sensitiveText := "成小王微笑着对毒品销售说，我认为台湾国的人有点意思"

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
