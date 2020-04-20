xs
=================

[![Travis Status for KarlTango/xs](https://travis-ci.org/KarlTango/xs.svg?branch=master&label=linux+build)](https://travis-ci.org/github/KarlTango/xs)
[![codecov for KarlTango/xs](https://codecov.io/gh/KarlTango/xs/branch/master/graph/badge.svg)](https://codecov.io/gh/KarlTango/xs)
[![Go Report Card](https://goreportcard.com/badge/github.com/KarlTango/xs)](https://goreportcard.com/report/github.com/KarlTango/xs)
[![GolangCI](https://golangci.com/badges/github.com/KarlTango/xs.svg)](https://golangci.com)
[![Release](https://img.shields.io/github/release/KarlTango/xs.svg?label=Release)](https://github.com/KarlTango/xs/releases)

[English](README.md)

安装
---

    go get github.com/KarlTango/xs

依赖
---

* github.com/tealeg/xlsx

目录
---

1. [例子](#examples)
	1. [Unmarshal](#examples_unmarshal)
   	2. [Marshal](#examples_marshal)

例子 <a name="examples"></a>
-------------------------

### Unmarshal <a name="examples_unmarshal"></a>
**如果你有一个像这样的excel:**

![Excel file](example/excel.png "Excel")
(样例文件放在 example 文件夹下).

**首先你可以定义一个结构体， 其中xs tag 对应着excel的表头**
```golang
type User struct {
	Username string `xs:"Username"`
	Age      int    `xs:"Age"`
	Sex      string `xs:"Sex"`
}
```

**然后你可以使用 UnmarshalFromFile方法将一个指定路径的excel文件反序列化为go struct**
```golang
var users = make([]User, 0)
filePath := "example/test.xlsx"
err := UnmarshalFromFile(filePath, &users)
```
**或者使用 Unmarshal方法将一个字节数组反序列化为go struct**
```golang
var users = make([]User, 0)
filePath := "example/test.xlsx"
body, _ := ioutil.ReadFile(filePath)
err = Unmarshal(body, &users)
```


### Marshal <a name="examples_marshal"></a>
**同样的，你可以使用 Marshal方法来将你的结构体数组写入到一个excel文件中**
#### 你可以传递多个结构体数组，他们会被同时写入一个excel的不同sheet中
```golang
type User struct {
	Username string `xs:"Username"`
	Age      int    `xs:"Age"`
	Sex      string `xs:"Sex"`
}

type User2 struct {
	Username string  `xs:"Username"`
	Salary   float64 `xs:"Salary"`
}

users := []User{{"Karl", 25, "Male"}, {"Ann", 18, "Female"}}
users2 := []User2{{"Karl", 8000.00}, {"Ann", 9999.50}}

//marshal
body, err := Marshal(users, users2)
if err != nil {
    //TODO: process err
}

//write file
err = ioutil.WriteFile("user.xlsx", body, 06666)
if err != nil {
    //TODO: process err
}
```

