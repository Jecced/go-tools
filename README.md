<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [go-tools](#go-tools)
- [通用系](#%E9%80%9A%E7%94%A8%E7%B3%BB)
  - [文件系统路径分隔符](#%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E8%B7%AF%E5%BE%84%E5%88%86%E9%9A%94%E7%AC%A6)
- [组件类](#%E7%BB%84%E4%BB%B6%E7%B1%BB)
  - [set](#set)
- [雪花算法](#%E9%9B%AA%E8%8A%B1%E7%AE%97%E6%B3%95)
- [伪随机数](#%E4%BC%AA%E9%9A%8F%E6%9C%BA%E6%95%B0)
- [图片操作相关](#%E5%9B%BE%E7%89%87%E6%93%8D%E4%BD%9C%E7%9B%B8%E5%85%B3)
- [谷歌翻译(英文转中文)](#%E8%B0%B7%E6%AD%8C%E7%BF%BB%E8%AF%91%E8%8B%B1%E6%96%87%E8%BD%AC%E4%B8%AD%E6%96%87)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# go-tools
个人go语言开发工具集

# 通用系 
## 文件系统路径分隔符 
```go
ak.PS
// 用来快速替代
string(os.PathSeparator)
```

# 组件类
## set
```go
// 获取一个非线程安全的set
var set Set = comp.NewSet(...items)

// 获取一个线程安全的set
var set Set = comp.NewCSet(...items)

set.Add(item node) bool
set.Remove(item node) bool
set.Has(item node) bool
set.Size() int
set.Clear()
set.IsEmpty() bool
set.List() []node

```

# 雪花算法
```go
// workerId 工作ID (0~31)
// datacenterId 数据中心ID (0~31)
worker, err := snowflake.CreateWorker(0, 0)
var id int64 = worker.NextId()
// 将十进制数字转化为二进制字符串
var sid string = snowflake.ConvertToBin(id)
```

# 伪随机数
```go
rand := randutil.Random(1000)

rand.GetSeed() //1000
rand.SetSeed(1000)

rand.Next(100, 200)
rand.NextInt(100, 200)
rand.NextInt32(100, 200)
rand.NextInt64(100, 200)
rand.NextBool()
```

# 图片操作相关
```go
// 加载一个图像
imgutil.LoadImage(path string) (img image.Image, err error)
// 保存一个图像
imgutil.SaveImage(p string, src image.Image) error

// img 图像, 顺时针旋转90°
imgutil.RotationRight(m image.Image) *image.RGBA
// img 图像, 逆时针旋转90°
imgutil.RotationLeft(m image.Image) *image.RGBA
// img 图像, 180°旋转
imgutil.Rotation180(m image.Image) *image.RGBA
// img 图像, 左右镜像翻转
imgutil.FlipMirror(m image.Image) *image.RGBA
// img 图像, 上下垂直翻转
imgutil.FlipVertical(m image.Image) *image.RGBA
// img 文件, 顺时针旋转90°
imgutil.RotationImgRight(in, out string) error
// img 文件, 逆时针旋转90°
imgutil.RotationImgLeft(in, out string) error
// img 文件, 180°旋转
imgutil.RotationImg180(in, out string) error
// img 文件 镜像翻转
imgutil.FlipImgMirror(in, out string) error
// img 文件 上下垂直翻转
imgutil.FlipImgVertical(in, out string) error

// 将一个文件夹内的内容都进行Trim
imgutil.TrimBlankDir(dir, suffix string) error
// 将一个图片文件Trim空白边
imgutil.TrimBlankFile(in, out string) error 
// 将一个图片对象Trim空白边
imgutil.TrimBlankImg(src image.Image) (image.Image, error)

// 图片裁剪
imgutil.Trimming(in, out string, x, y, w, h int) error
// 创建一个指定宽高的图片
imgutil.CreatPng(width, height int) *image.RGBA
// 混合两个图片
imgutil.MixImg(src *image.RGBA, dist image.Image, x, y int)

// Resample返回图像切片r (m)的重新采样副本。
// 返回的图像宽度为w，高度为h。
imgutil.Resize(m image.Image, w, h int) *image.RGBA
```

# 谷歌翻译(英文转中文)
```go
got := gts.TranslateEn2Cn("今天天气还不错")
fmt.Println(got)
// console
// The weather is pretty good today
```
