<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [网络请求使用库](#%E7%BD%91%E7%BB%9C%E8%AF%B7%E6%B1%82%E4%BD%BF%E7%94%A8%E5%BA%93)
- [Usage | 用法](#usage--%E7%94%A8%E6%B3%95)
  - [Simple Case | 简单案例](#simple-case--%E7%AE%80%E5%8D%95%E6%A1%88%E4%BE%8B)
  - [WriteToFile | 写入到文件](#writetofile--%E5%86%99%E5%85%A5%E5%88%B0%E6%96%87%E4%BB%B6)
  - [Parameters | 设置请求参数](#parameters--%E8%AE%BE%E7%BD%AE%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  - [Set Headers | 设置请求头](#set-headers--%E8%AE%BE%E7%BD%AE%E8%AF%B7%E6%B1%82%E5%A4%B4)
  - [Timeout | 设置超时时间](#timeout--%E8%AE%BE%E7%BD%AE%E8%B6%85%E6%97%B6%E6%97%B6%E9%97%B4)
  - [Proxy | 代理](#proxy--%E4%BB%A3%E7%90%86)
  - [BasicAuth | 基础认证](#basicauth--%E5%9F%BA%E7%A1%80%E8%AE%A4%E8%AF%81)
  - [Cookie | Cookie](#cookie--cookie)
  - [Session | Session](#session--session)
    - [Session.BaseAuth | 基于 Session 的 基础认证](#sessionbaseauth--%E5%9F%BA%E4%BA%8E-session-%E7%9A%84-%E5%9F%BA%E7%A1%80%E8%AE%A4%E8%AF%81)
    - [Session.Proxy | 基于 session 的代理](#sessionproxy--%E5%9F%BA%E4%BA%8E-session-%E7%9A%84%E4%BB%A3%E7%90%86)
    - [Session.Header | 基于 session 的请求头](#sessionheader--%E5%9F%BA%E4%BA%8E-session-%E7%9A%84%E8%AF%B7%E6%B1%82%E5%A4%B4)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 网络请求使用库

让 `golang` 更简单的发起 `http` 请求

链式的请求规则

Convenient http client for go.

# Usage | 用法

## Simple Case | 简单案例

一个简单的http请求示例, 示例执行http请求, 并将响应读取为字符串 

One simple http request example that do http get request and read response as string:
```go
resp, err := https.Get("http://www.baidu.com/").
            Send().
            ReadText()
fmt.Println(resp)
```

## WriteToFile | 写入到文件

将响应请求保存到本地文件中, 例如保存一个图片

Save the response request to a local file, such as an image
```go
err := https.Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").
		Send().
		SetTimeOut(30_000).
		WriteToFile("/Users/ankang/develop/test/1.png")
```

## Parameters | 设置请求参数

Pass parameters in urls using params method:
```go
resp, err := https.Get(url).
            AddParam("key1", "value1").
            AddParam("key2", "value2").
            AddParam("key3", "value3").
            Send().
            ReadText()
fmt.Println(resp)
```
OR
```go
resp, err := https.Get(url).
            AddParams(map[string]string{
                "key1": "value1",
                "key2": "value2",
                "key3": "value3",
            }).
            Send().
            ReadText()
fmt.Println(resp)
```

## Set Headers | 设置请求头

Http request headers can be set by headers method:
```go
resp, err := https.Get(url).
            AddHeader("key1", "value1").
            AddHeader("key2", "value2").
            AddHeader("key3", "value3").
            Send().
            ReadText()
fmt.Println(resp)
```
OR
```go
resp, err := https.Get(url).
            AddHeaders(map[string]string{
                "key1": "value1",
                "key2": "value2",
                "key3": "value3",
            }).
            Send().
            ReadText()
fmt.Println(resp)
```

## Timeout | 设置超时时间

你可以设置连接超时时间, 和响应请求的超时时间

You can set connection connect timeout, and socket read/write timeout value
```go
https.Get(url).
    SetConnTimeOut(30_000).
    SetRespTimeOut(30_000).
    Send().
    ReadText()
```

你也可以同时修改他们

You can also change their values at the same time
```go
https.Get(url).
    SetTimeOut(30_000).
    Send().
    ReadText()
```

## Proxy | 代理

通过代理方法设置代理

Set proxy by proxy method
```go
resp, err := https.Get("http://www.google.com").
            Proxy("127.0.0.1:1081").
            Send().
            ReadText()
fmt.Println(resp)
```

## BasicAuth | 基础认证

```go
resp := https.Get(url).
    BasicAuth(user_name, password).
    Send().
    ReadText()
fmt.Println(resp)
```

## Cookie | Cookie

```go
resp, err := https.Get(url).
    AddCookie("key1", "value1").
    AddCookies(map[string]string{
        "key2", "value2",
        "key3", "value3",
        "key4", "value4",
    }).
    Send().
    ReadText()
fmt.Println(resp)
```

## Session | Session

会话为您维护cookie，在需要登录或其他情况时很有用。会话的用法与请求相同。

Session maintains cookies, useful when need login or other situations. Session have the same usage as Requests.

```go
s := https.Session()
get, err := s.Get(url1).Send().ReadText()
post, err := s.Post(url2).Send().ReadText()
fmt.Println(get)
fmt.Println(post)
```

### Session.BaseAuth | 基于 Session 的 基础认证

所有由session发起的请求都会用这个基础认证信息

add session method new func

```go
session, err := https.Session()
session.BasicAuth("user", "password")
session.Get(url).Send().ReadText()
```

### Session.Proxy | 基于 session 的代理

所有由session发起的请求都会用这个代理信息

add session method new func to use proxy

```go
session := https.Session().Proxy("127.0.0.1:1081")
session.Get(url1).Send().ReadText()
session.Get(url2).Send().ReadText()
```

### Session.Header | 基于 session 的请求头

所有由session发起的请求都会用这些请求头

add session method new func to use header

```go
session := https.Session().
    AddHeader("header1", "value1").
    AddHeader("header2", "value2").
    AddHeaders(map[string]string{
        "key1": "value1",
        "key2": "value2",
        "key3": "value3",
    })
session.Get(url1).Send().ReadText()
session.Get(url2).Send().ReadText()
```
