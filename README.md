xs
=================

[![Travis Status for KarlTango/xs](https://travis-ci.org/KarlTango/xs.svg?branch=master&label=linux+build)](https://travis-ci.org/github/KarlTango/xs)
[![codecov for KarlTango/xs](https://codecov.io/gh/KarlTango/xs/branch/master/graph/badge.svg)](https://codecov.io/gh/KarlTango/xs)
[![Go Report Card](https://goreportcard.com/badge/github.com/KarlTango/xs)](https://goreportcard.com/report/github.com/KarlTango/xs)

Installation
------------

    go get github.com/KarlTango/xs

Requirements
-------------

* github.com/tealeg/xlsx

Index
------

1. [Examples](#examples)
	1. [Unmarshal](#examples_unmarshal)
   	2. [Marshal](#examples_marshal)

Examples <a name="examples"></a>
--------

### Unmarshal <a name="examples_unmarshal"></a>
**If you wan to get message from a excel file like this:**

![Excel file](example/excel.png "Excel")
(Excel could be found in example folder).

**You could define struct with xs tag to map struct key excel column name**
```golang
type User struct {
	Username string `xs:"Username"`
	Age      int    `xs:"Age"`
	Sex      string `xs:"Sex"`
}
```

**and then you can use UnmarshalFromFile**
```golang
var users = make([]User, 0)
filePath := "example/test.xlsx"
err := UnmarshalFromFile(filePath, &users)
```
**or use Unmarshal**
```golang
var users = make([]User, 0)
filePath := "example/test.xlsx"
body, _ := ioutil.ReadFile(filePath)
err = Unmarshal(body, &users)
```


### Marshal <a name="examples_marshal"></a>
**The other way round, if you want to write a excel from go struct slice, you can use Marshal**
#### Mutiple sheet
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

