<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [go-tools | 使用方式](#go-tools--%E4%BD%BF%E7%94%A8%E6%96%B9%E5%BC%8F)
- [通用系](#%E9%80%9A%E7%94%A8%E7%B3%BB)
  - [文件系统路径分隔符](#%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E8%B7%AF%E5%BE%84%E5%88%86%E9%9A%94%E7%AC%A6)
- [组件类](#%E7%BB%84%E4%BB%B6%E7%B1%BB)
  - [set](#set)
- [雪花算法](#%E9%9B%AA%E8%8A%B1%E7%AE%97%E6%B3%95)
- [伪随机数](#%E4%BC%AA%E9%9A%8F%E6%9C%BA%E6%95%B0)
- [图片操作相关](#%E5%9B%BE%E7%89%87%E6%93%8D%E4%BD%9C%E7%9B%B8%E5%85%B3)
  - [加载保存一张图片](#%E5%8A%A0%E8%BD%BD%E4%BF%9D%E5%AD%98%E4%B8%80%E5%BC%A0%E5%9B%BE%E7%89%87)
  - [图片的旋转](#%E5%9B%BE%E7%89%87%E7%9A%84%E6%97%8B%E8%BD%AC)
  - [去除图片四周空白透明](#%E5%8E%BB%E9%99%A4%E5%9B%BE%E7%89%87%E5%9B%9B%E5%91%A8%E7%A9%BA%E7%99%BD%E9%80%8F%E6%98%8E)
  - [图片缩放](#%E5%9B%BE%E7%89%87%E7%BC%A9%E6%94%BE)
  - [图片工具](#%E5%9B%BE%E7%89%87%E5%B7%A5%E5%85%B7)
  - [gzip](#gzip)
- [谷歌翻译(英文转中文)](#%E8%B0%B7%E6%AD%8C%E7%BF%BB%E8%AF%91%E8%8B%B1%E6%96%87%E8%BD%AC%E4%B8%AD%E6%96%87)
- [网络请求工具](#%E7%BD%91%E7%BB%9C%E8%AF%B7%E6%B1%82%E5%B7%A5%E5%85%B7)
- [file 操作](#file-%E6%93%8D%E4%BD%9C)
- [string 操作](#string-%E6%93%8D%E4%BD%9C)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# go-tools | 使用方式
```shell
go get github.com/Jecced/rs
```

# 通用系 
## 文件系统路径分隔符 
```go
ak.PS
// 用来快速替代
string(os.PathSeparator)

// 获取md5
GetMd5(data *[]byte) string 
// 生成base64
EncodeBase64(data *[]byte) string 
// 解析base64
DecodeBase64(text string) ([]byte, error) 
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
## 加载保存一张图片
```go
// 加载一个图像
imgutil.LoadImage(path string) (img image.Image, err error)
// 保存一个图像
imgutil.SaveImage(p string, src image.Image) error
```
## 图片的旋转
```go
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
```
## 去除图片四周空白透明
```go
// 将一个文件夹内的内容都进行Trim
imgutil.TrimBlankDir(dir, suffix string) error
// 将一个图片文件Trim空白边
imgutil.TrimBlankFile(in, out string) error 
// 将一个图片对象Trim空白边
imgutil.TrimBlankImg(src image.Image) (image.Image, error)
```
## 图片缩放
```go
// Resample返回图像切片r (m)的重新采样副本。
// 返回的图像宽度为w，高度为h。
imgutil.Resize(m image.Image, w, h int) *image.RGBA
```
## 图片工具
```go
// 图片裁剪
imgutil.Trimming(in, out string, x, y, w, h int) error
// 创建一个指定宽高的图片
imgutil.CreatPng(width, height int) *image.RGBA
// 混合两个图片
imgutil.MixImg(src *image.RGBA, dist image.Image, x, y int)
```

## gzip
```go
gziputil.ZIP([]byte, io.Writer) (int, error)
gziputil.UNZIP(io.Reader) ([]byte, error)
```

# 谷歌翻译(英文转中文)
```go
got := translate.GoogleTranslate("今天天气还不错")
fmt.Println(got)
// console
// The weather is pretty good today
```

# 网络请求工具
**接口过多, 请访问独立文档**  *[请求工具文档](https://github.com/Jecced/go-tools/blob/master/README_HTTPS.md)*

# file 操作
```go
// 根据路径创建文件夹
fileutil.MkdirAll(path string) error
// 创建一个文件的父目录
fileutil.MkdirParent(path string) error
// 获取某个文件夹下所有指定后缀的文件
fileutil.GetFilesBySuffix(dirPath string, suffix string) (files []string, err error)
// 文件拷贝
fileutil.FileCopy(src, dist string) (err error)
// 目录拷贝
fileutil.DirCopy(src, dist string) error
// 判断一个路径是否存在
fileutil.PathExists(path string) bool
// 获取一个路径的父目录地址
fileutil.GetParentDir(path string) string
// 路径格式化, 标准化一个路径到当前系统规范
fileutil.PathFormat(path string) string
// 获取目录下所有文件类型
fileutil.FindAllFileTypes(dir string) (types []string)
// 写入文本到指定文件
fileutil.WriteText(text, dist string) error
// 写入数据到指定文件
fileutil.WriteData(data []byte, dist string) error
// 将一个文件读取成字符串返回
fileutil.ReadText(file string) (string, error)
// 读取一个文件的 byte 二进制
fileutil.ReadBytes(file string) ([]byte, error)
// 清空一个目录的所有内容
fileutil.ClearDir(dir string)
// 可以删除一个文件, 空文件夹
fileutil.RemoveFile(file string) error
// 自内向外删除所有空文件夹, 如果文件是.DS_Store的话, 也会一起删除
fileutil.DelEmptyDir(dir string) (bool, error)
```


# string 操作
```go
// 判断字符串是否包含中文
strutil.IsChineseChar(str string) bool

// 下划线转驼峰
strutil.UnderscoreToUpperCamelCase(s string) string

// 下划线转小写驼峰
strutil.UnderscoreToLowerCamelCase(s string) string

// 驼峰转下划线
strutil.CamelCaseToUnderscore(s string) string

// 前后缀匹配出第一个[正则方式]
strutil.FindMatchFirst(str, s, e string, fix bool) string

// 前后缀匹配[正则方式]
strutil.FindMatch(str, s, e string, fix bool) []string

// 基于 strings.Index 实现的前后缀匹配查找第一个
strutil.MatchStringFirst(text, prefix, suffix string, fix bool) string

// 基于 strings.Index 实现的前后缀匹配查找
strutil.MatchString(text, prefix, suffix string, fix bool) []string

// 完善 strings.Index, 多加入索引位置参数
strutil.IndexOf(text, substr string, index int) int

// 插入文本, 插入的内容 @insert 会放在原始文本 @text 中 @template 的前面
// 在 @insert 中搜索 @template 的位置 @st
// 在 @st 位置 前 插入 @insert 的字符串文本内容
// @param text 		原始文本
// @param insert 	插入的内容
// @param template 	查询字符串
strutil.InsertString(text *string, insert, template string)

// 转义\u00e9文字
// 转义\xE9\x80文字
strutil.Decode(text string) (string, error)
```

# LICENSE

    MIT License
    
    Copyright (c) 2020 Jecced
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.