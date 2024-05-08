# go-sensitive-word

敏感词（敏感词/违禁词/违法词/脏词）检测工具，基于 DFA 算法实现的高性能 Go 敏感词过滤工具框架。

## 快速接入

**安装**
```bash
go get -u e.coding.net/zmexing/zx/go-sensitive-word@latest
```

**使用**
```go

```

## 常见问题

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