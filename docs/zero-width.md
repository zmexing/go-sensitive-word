# 零宽字符攻击

在Go语言中，防止零宽字符攻击可以通过一些方法来实现。零宽字符攻击通常是通过在字符串中插入隐藏的零宽字符来欺骗系统，例如在认证字符串中插入零宽空格来绕过验证等。以下是一些防止零宽字符攻击的方法：

- 验证输入字符串：在接受用户输入的地方，进行输入验证，确保字符串中不包含任何不必要的或可疑的字符。可以使用正则表达式或简单的字符检查来过滤或拒绝包含零宽字符的字符串。
- 规范化字符串：在处理输入字符串之前，可以使用Unicode规范化函数将字符串标准化为特定形式，以移除任何不必要的或隐藏的字符。
- 使用字节处理而不是字符串：在某些情况下，将字符串视为字节序列可能更安全。通过将字符串转换为字节切片并以字节为单位进行处理，可以避免字符串中的隐藏字符问题。

```go
// IsZeroWidth 是否存在零宽字符
func IsZeroWidth(s string) bool {
	re := regexp.MustCompile(`[\p{Cf}\\u200B]`)
	return re.MatchString(s)
}

// RemoveZeroWidth 移除零宽字符
func RemoveZeroWidth(s string) string {
    re := regexp.MustCompile(`[\p{Cf}\\u200B]`)
    return re.ReplaceAllString(s, "")
}
```

```go
import (
    "fmt"
    "testing"
)

const (
	str1     = "马斯克"
	str2Zero = "马​‌‍斯​‌‍克​‌‍"
	str3Zero = "马\\u200B斯\\u200B克"
)

func TestIsZeroWidth(t *testing.T) {
	res1 := IsZeroWidth(str1)
	fmt.Println("res1 结果：", res1) // false

	res2 := IsZeroWidth(str2Zero)
	fmt.Println("res2 结果：", res2) // true

	res3 := IsZeroWidth(str3Zero)
	fmt.Println("res3 结果：", res3) // true
}

func TestRemoveIsZeroWidth(t *testing.T) {
    res1 := RemoveZeroWidth(str1)
    fmt.Println("res1 结果：", res1) // 马斯克
    
    res2 := RemoveZeroWidth(str3Zero)
    fmt.Println("res2 结果：", res2) // 马斯克
    
    res3 := res1 == res2
    fmt.Println("res1 和 res2 比较的结果：", res3)
}
```