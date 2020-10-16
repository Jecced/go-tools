package strutil

import (
	"regexp"
	"strings"
	"unicode"
)

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
	rex := regexp.MustCompile(s + `([^` + e + `]+)` + e)
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
