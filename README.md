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


// 输出结果
//res1: true
//res2: 成小王
//res3: [成小王 毒品 毒品销售 台湾国]
//res4: map[台湾国:1 成小王:1 毒品:1 毒品销售:1]
//res5: ***微笑着对****说，我认为***的人有点意思
//res6: 微笑着对销售说，我认为的人有点意思 
```

## 更多特性

### 字符串检测

```go
// HasEmail 判断字符串中是否存在邮箱地址
func HasEmail(s string) bool

// MaskEmail 将字符串中存在的邮箱地址替换成 "*"
func MaskEmail(s string) string

// HasURL 判断字符串中是否存在网址
func HasURL(s string) bool

// MaskURL 将字符串中存在的网址替换成 "*"
func MaskURL(s string) string

// HasDigit 判断字符串中是否存在指定个数的数字（大于等于该数字）
func HasDigit(s string, count int) bool

// MaskDigit 将字符串中存在的数字替换成 "*"
func MaskDigit(s string) string

// HasWechatID 判断字符串中是否存在微信号
func HasWechatID(s string) bool

// MaskWechatID 将字符串中存在的微信号替换成 "*"
func MaskWechatID(s string) string
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

## 参考资料
- 基于Java DFA实现的敏感词过滤：https://github.com/houbb/sensitive-word
- unicode字词的神奇组合：https://www.zhihu.com/question/30873035
- unicode违规技巧：https://zhuanlan.zhihu.com/p/545309061
- unicode视觉欺骗：https://zhuanlan.zhihu.com/p/611904676
- unicode字符列表：https://symbl.cc/en/unicode-table
- 汉字结构描述字符：https://zh.wikipedia.org/wiki/%E8%A1%A8%E6%84%8F%E6%96%87%E5%AD%97%E6%8F%8F%E8%BF%B0%E5%AD%97%E7%AC%A6