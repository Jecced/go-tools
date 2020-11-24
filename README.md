<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [go-tools](#go-tools)
- [通用系](#%E9%80%9A%E7%94%A8%E7%B3%BB)
  - [文件系统路径分隔符](#%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E8%B7%AF%E5%BE%84%E5%88%86%E9%9A%94%E7%AC%A6)
- [组件类](#%E7%BB%84%E4%BB%B6%E7%B1%BB)
  - [set](#set)
- [雪花算法](#%E9%9B%AA%E8%8A%B1%E7%AE%97%E6%B3%95)
- [伪随机数](#%E4%BC%AA%E9%9A%8F%E6%9C%BA%E6%95%B0)
- [谷歌翻译(英文转中文)](#%E8%B0%B7%E6%AD%8C%E7%BF%BB%E8%AF%91%E8%8B%B1%E6%96%87%E8%BD%AC%E4%B8%AD%E6%96%87)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# go-tools
个人go语言开发工具集

# 通用系 
## 文件系统路径分隔符 
```
ak.PS
# 用来快速替代
string(os.PathSeparator)
```

# 组件类
## set
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

# 雪花算法
```
// workerId 工作ID (0~31)
// datacenterId 数据中心ID (0~31)
worker, err := snowflake.createWorker(0, 0)
worker.nextId()
```

# 伪随机数
```
rand := randutil.Random(1000)

rand.GetSeed() //1000
rand.SetSeed(1000)

rand.Next(100, 200)
rand.NextInt(100, 200)
rand.NextInt32(100, 200)
rand.NextInt64(100, 200)
rand.NextBool()
```



# 谷歌翻译(英文转中文)
```
got := gts.TranslateEn2Cn("今天天气还不错")
fmt.Println(got)
// console
// The weather is pretty good today
```
