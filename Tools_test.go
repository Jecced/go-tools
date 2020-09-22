package go_tools

import (
	"fmt"
	"go-tools/src"
	"go-tools/src/fileutil"
	"regexp"
	"strings"
	"sync"
	"testing"
)

func TestTranslateEn2Cn(t *testing.T) {
	got := gts.TranslateEn2Cn("今天天气还不错")
	fmt.Println(got)
}

var mutex sync.Mutex
var cache = make(map[string]string)
var src = ""
var out = "/Users/ankang/百度云/DoodleGodRocketScientist/static/i18n/en-US/en-US2.json"
var all = 0

func TestLoadCn(t *testing.T) {
	quotes, _ := fileutil.ReadText("/Users/ankang/百度云/DoodleGodRocketScientist/static/i18n/en-US/en-US2.json")
	src = quotes
	quotes = strings.ReplaceAll(quotes, ":", ":\n")

	compile := regexp.MustCompile(`".+"`)
	params := compile.FindAll([]byte(quotes), 100000000)

	for _, param := range params {
		if strings.Index(string(param), "_$_") != -1 {
			continue
		}
		v := strings.ReplaceAll(string(param), "\\n", " ")
		v = strings.ReplaceAll(v, "_", " ")
		cache[string(param)] = v
	}
	all = len(cache)
	go display()
	go display()
	select {}
}

func display() {
	for {
		k, v := getValue()
		if k == "" && v == "" {
			fmt.Println("线程停止...")
			break
		}
		got := gts.TranslateEn2Cn(v[1 : len(v)-1])
		if "" == got {
			continue
		}
		if 0 == len(v) {
			continue
		}
		out := fmt.Sprintf("\"%s_$_%s\"", v[1:len(v)-1], got)
		replace(k, out)
		//replace(k, "\""+v+"_#$#_"+got+"\"")
		if len(got) < 200 {
			save()
		}
		fmt.Printf("======进度:%d/%d\n%s\n%s\n", all-len(cache), all, v, got)
	}
}

func getValue() (k, v string) {
	mutex.Lock()
	kk, vv := "", ""
	for k1, v1 := range cache {
		delete(cache, k1)
		kk, vv = k1, v1
		break
	}
	mutex.Unlock()
	return kk, vv
}

func replace(a, d string) {
	mutex.Lock()
	src = strings.ReplaceAll(src, a, d)
	mutex.Unlock()
}

func save() {
	mutex.Lock()
	fileutil.WriteText(src, out)
	mutex.Unlock()
}

func Test000(t *testing.T) {
	got := gts.TranslateEn2Cn("Baba! Sparkle sparkle!")
	fmt.Println(got)
}
