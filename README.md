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
comp.Set(...items)

// 获取一个线程安全的set
comp.CSet(...items)

set.Add(item key)
set.Remove(item key)
set.Has(item key)
set.Size(item key)
set.Clear(item key)
set.IsEmpty(item key)
set.List(item key)

```



## 谷歌翻译(英文转中文)
```
got := gts.TranslateEn2Cn("今天天气还不错")
fmt.Println(got)
// console
// The weather is pretty good today
```
