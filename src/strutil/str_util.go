package strutil

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 判断字符串是否包含中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) ||
			regexp.MustCompile(
				"[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]",
			).MatchString(string(r)) {
			return true
		}
	}
	return false
}

// 下划线转驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s1 := UnderscoreToUpperCamelCase(s)
	if s1 == "" {
		return s
	}
	return string(unicode.ToLower(rune(s1[0]))) + s1[1:]
}

// 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// 前后缀匹配出第一个
func FindMatchFirst(str, s, e string, fix bool) string {
	resp := FindMatch(str, s, e, fix)
	if 0 == len(resp) {
		return ""
	}
	return resp[0]
}

// 前后缀匹配
func FindMatch(str, s, e string, fix bool) []string {
	ex := s + `([^(` + e + `|` + s + `)]+)` + e
	rex := regexp.MustCompile(ex)
	//rex := regexp.MustCompile(`name="([^"]+)"`)
	out := rex.FindAllStringSubmatch(str, -1)

	var resp []string

	for _, i := range out {
		if fix {
			resp = append(resp, i[0])
		} else {
			resp = append(resp, i[1])
		}
	}
	return resp
}

// 基于 strings.Index 实现的前后缀匹配查找第一个
func MatchStringFirst(text, prefix, suffix string, fix bool) string {
	resp := MatchString(text, prefix, suffix, fix)
	if 0 == len(resp) {
		return ""
	}
	return resp[0]
}

// 基于 strings.Index 实现的前后缀匹配查找
func MatchString(text, prefix, suffix string, fix bool) []string {
	st, ed := -1, -1
	prefixLen := len(prefix)
	suffixLen := len(suffix)
	out := make([]string, 0, 0)
	for st = IndexOf(text, prefix, st); st != -1; st = IndexOf(text, prefix, st) {
		ed = IndexOf(text, suffix, st+prefixLen)
		if ed == -1 || ed == st {
			break
		}
		t := ""
		if fix {
			t = text[st : ed+suffixLen]
		} else {
			t = text[st+prefixLen : ed]
		}
		st = ed + suffixLen
		out = append(out, t)
	}
	return out
}

// 完善 strings.Index, 多加入索引位置参数
func IndexOf(text, substr string, index int) int {
	if index < 0 {
		index = 0
	}
	i := strings.Index(text[index:], substr)
	if -1 == i {
		return -1
	}
	return i + index
}

// 插入文本, 插入的内容 @insert 会放在原始文本 @text 中 @template 的前面
// 在 @insert 中搜索 @template 的位置 @st
// 在 @st 位置 前 插入 @insert 的字符串文本内容
// @param text 		原始文本
// @param insert 	插入的内容
// @param template 	查询字符串
func InsertString(text *string, insert, template string) {
	st := strings.Index(*text, template)
	if -1 == st {
		return
	}
	*text = (*text)[:st] + insert + (*text)[st:]
}

// 转义\u00e9文字
// 转义\xE9\x80文字
func Decode(text string) (string, error) {
	return strconv.Unquote("\"" + text + "\"")
}
