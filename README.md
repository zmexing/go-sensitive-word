# go-sensitive-word

敏感词（敏感词/违禁词/违法词/脏词）检测工具，基于 DFA 算法实现的高性能 Go 敏感词过滤工具框架。

## 快速接入

**安装**
```bash
go get -u e.coding.net/zmexing/zx/go-sensitive-word@latest
```

**使用**
```go
package main

import (
	sensitive "e.coding.net/zmexing/zx/go-sensitive-word"
	"fmt"
	"log"
)

func main() {
	filter := sensitive.NewFilter(
		sensitive.StoreOption{Type: sensitive.StoreMemory},
		sensitive.FilterOption{Type: sensitive.FilterDfa},
	)

	// 加载敏感词库
	err := filter.Store.LoadDictPath("./your_path/dict.txt")
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


// 输出结果
// res1: true
// res2: 毒品
// res3: [毒品 毒品销售 台湾国]
// res4: map[台湾国:1 毒品:1 毒品销售:1]
// res5: 李先生微笑着对****说，我认为***的人有点意思
// res6: 李先生微笑着对销售说，我认为的人有点意思 
```

## 常见问题

- [Unicode相似字符攻击](./docs/unicode.md)
- [零宽字符攻击](docs/zero-width.md)

## 性能压测

**硬件环境**
- 处理器：Apple M2
- 基带：内存：16 GB 类型：LPDDR5
- 系统类型：MacOS Ventura 版本13.5 (22G74)

**压测结果**

| 方法 | Go耗时                | Java耗时           | 备注                   |
|---|---------------------|------------------|----------------------|
| IsSensitive  | 267.6 ns，约 452W QPS | 240 ms，约 10W QPS | 大约是Java性能的 898,876 倍 |
| Replace | 587.5 ns，约 202W QPS | 447 ms，约 10W QPS | 大约是Java性能的 761,702 倍 |