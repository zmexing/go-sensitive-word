# 性能压测

**硬件环境**
- 处理器：Apple M2
- 基带：内存：16 GB 类型：LPDDR5
- 系统类型：MacOS Ventura 版本13.5 (22G74)

**压测结果**

| 方法 | Go耗时              | Java耗时             | 备注               |
|---|-------------------|--------------------|------------------|
| IsSensitive  | 267 ns，约 452W QPS | 2724 ns，约 200W QPS | 大约是Java性能的 10 倍  |
| Replace | 587 ns，约 202W QPS | 2865 ns，约 200W QPS | 大约是Java性能的 4.8 倍 |

## 压测代码

**Go**

```go
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
```

**Java**

```java
package com.github.houbb.sensitive.word.benchmark;

import com.github.houbb.heaven.util.util.RandomUtil;
import com.github.houbb.sensitive.word.bs.SensitiveWordBs;
import com.github.houbb.sensitive.word.core.SensitiveWordHelper;
import org.junit.Ignore;
import org.junit.Test;

@Ignore
public class BenchmarkTimesTest {
    
    @Test
    public void onlyContainsTest() {
        SensitiveWordBs sensitiveWordBs = SensitiveWordBs.newInstance()
                .enableWordCheck(true)
                .enableNumCheck(false)
                .enableUrlCheck(false)
                .enableEmailCheck(false)
                .ignoreRepeat(false)
                .ignoreCase(false)
                .ignoreNumStyle(false)
                .ignoreChineseStyle(false)
                .ignoreEnglishStyle(false)
                .ignoreWidth(false)
                .init();

        String randomText = "小明微笑着对毒品销售说，我认为台湾国的人有点意思";

        long start = System.nanoTime();
        for(int i = 0; i < 2000000; i++) {
            sensitiveWordBs.contains(randomText);
        }
        long end = System.nanoTime();
        System.out.println("------------------ COST: " + (end-start)/2000000);
    }

    @Test
    public void onlyReplaceTest() {
        SensitiveWordBs sensitiveWordBs = SensitiveWordBs.newInstance()
                .enableWordCheck(true)
                .enableNumCheck(false)
                .enableUrlCheck(false)
                .enableEmailCheck(false)
                .ignoreRepeat(false)
                .ignoreCase(false)
                .ignoreNumStyle(false)
                .ignoreChineseStyle(false)
                .ignoreEnglishStyle(false)
                .ignoreWidth(false)
                .init();

        String randomText = "小明微笑着对毒品销售说，我认为台湾国的人有点意思";

        long start = System.nanoTime();
        for(int i = 0; i < 2000000; i++) {
            sensitiveWordBs.replace(randomText);
        }
        long end = System.nanoTime();
        System.out.println("------------------ COST: " + (end-start)/2000000);
    }
}
```