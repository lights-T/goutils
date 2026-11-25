package xstrings

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const letterBytes = "1234567890"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

//GetMd5String 生成32位md5字串
func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToStringByAscii(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func SplitData(input string, max int) (string, string) {
	// 使用";"作为分隔符分割输入字符串
	parts := strings.Split(input, ";")

	// 如果分割后的元素数量不超过100，则直接返回原字符串和空字符串
	if len(parts) <= max {
		return input, ""
	}

	// 构建A变量（前100个元素）
	aVariable := strings.Join(parts[:max], ";")

	// 构建B变量（剩余的元素）
	bVariable := strings.Join(parts[max:], ";")

	return aVariable, bVariable
}

// ParsePath 从形如 C:\ACM9\ACMSvc\ACMSvc9.exe 的路径中
// 提取盘符字母、以及前缀后的数字
func ParsePath(p string) (drive, num string, err error) {
	// 1. 盘符
	vol := filepath.VolumeName(p) // "C:"
	if vol == "" {
		return "", "", fmt.Errorf("无法识别盘符: %s", p)
	}
	drive = strings.TrimSuffix(vol, `:`)

	// 2. 取第一级目录名
	rest := p[len(vol):] // "\ACM9\ACMSvc\ACMSvc9.exe"
	rest = strings.TrimPrefix(rest, `\`)
	idx := strings.IndexByte(rest, '\\')
	first := rest
	if idx >= 0 {
		first = rest[:idx] // "ACM9"
	}

	// 3. 从末尾提取连续数字
	i := len(first)
	for i > 0 && first[i-1] >= '0' && first[i-1] <= '9' {
		i--
	}
	if i == len(first) {
		//没有数字，返回空字符串
		num = ""
		return
	}
	if i == 0 {
		return "", "", fmt.Errorf("目录名以数字开头，无前缀: %s", first)
	}

	//prefix = first[:i] // "ACM"
	num = first[i:] // "9" / "10" / "3"
	return
}
