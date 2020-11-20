   * [go-tools](#go-tools)
      * [通用系](#\xE9\x80\x9A\xE7\x94\xA8\xE7\xB3\xBB)
         * [ak.PS 文件系统路径分隔符](#akps-\xE6\x96\x87\xE4\xBB\xB6\xE7\xB3\xBB\xE7\xBB\x9F\xE8\xB7\xAF\xE5\xBE\x84\xE5\x88\x86\xE9\x9A\x94\xE7\xAC\xA6)
      * [组件类](#\xE7\xBB\x84\xE4\xBB\xB6\xE7\xB1\xBB)
         * [set](#set)
      * [谷歌翻译(英文转中文)](#\xE8\xB0\xB7\xE6\xAD\x8C\xE7\xBF\xBB\xE8\xAF\x91\xE8\x8B\xB1\xE6\x96\x87\xE8\xBD\xAC\xE4\xB8\xAD\xE6\x96\x87)



# go-tools
个人go语言开发工具集

## 通用系 
### `ak.PS` 文件系统路径分隔符 
```
# 用来快速替代
string(os.PathSeparator)
```

## 组件类
### set
```
// 获取一个非线程安全的set
var set Set = comp.Set(...items)

// 获取一个线程安全的set
var set Set = comp.CSet(...items)

set.Add(item node) bool
set.Remove(item node) bool
set.Has(item node) bool
set.Size() int
set.Clear()
set.IsEmpty() bool
set.List() []node

```



## 谷歌翻译(英文转中文)
```
got := gts.TranslateEn2Cn("今天天气还不错")
fmt.Println(got)
// console
// The weather is pretty good today
```
