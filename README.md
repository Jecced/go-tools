- [go-tools](#go-tools)
  - [通用系](#%E9%80%9A%E7%94%A8%E7%B3%BB)
    - [`ak.PS` 文件系统路径分隔符](#akps-%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E8%B7%AF%E5%BE%84%E5%88%86%E9%9A%94%E7%AC%A6)
  - [组件类](#%E7%BB%84%E4%BB%B6%E7%B1%BB)
    - [set](#set)
  - [谷歌翻译(英文转中文)](#%E8%B0%B7%E6%AD%8C%E7%BF%BB%E8%AF%91%E8%8B%B1%E6%96%87%E8%BD%AC%E4%B8%AD%E6%96%87)

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
