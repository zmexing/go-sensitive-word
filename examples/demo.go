package main

import (
	"fmt"
	sensitive "github.com/zmexing/go-sensitive-word"
	"log"
)

func main() {
	filter, err := sensitive.NewFilter(
		sensitive.StoreOption{Type: sensitive.StoreMemory}, // 基于内存
		sensitive.FilterOption{Type: sensitive.FilterDfa},  // 基于DFA算法
	)

	if err != nil {
		log.Fatalf("敏感词服务启动失败, err:%v", err)
		return
	}

	// 加载敏感词库
	err = filter.LoadDictEmbed(
		sensitive.DictCovid19,
		sensitive.DictOther,
		sensitive.DictReactionary,
		sensitive.DictViolence,
		sensitive.DictPeopleLife,
		sensitive.DictPornography,
		sensitive.DictAdditional,
		sensitive.DictCorruption,
		sensitive.DictTemporaryTencent,
	)
	if err != nil {
		log.Fatalf("加载词库发生了错误, err:%v", err)
		return
	}

	// 动态添加自定义敏感词
	err = filter.Store.AddWord("李世民", "秦始皇")
	if err != nil {
		log.Fatalf("添加敏感词发生了错误, err:%v", err)
		return
	}

	// 动态自定义敏感词白名单
	err = filter.Store.DelWord("武汉海鲜市场", "武汉")
	if err != nil {
		log.Fatalf("删除敏感词发生了错误, err:%v", err)
		return
	}

	sensitiveText := "李世民和老祖秦始皇的是忘年交，他们两个相约一起去武汉海鲜市场玩耍"

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
