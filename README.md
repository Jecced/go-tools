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
