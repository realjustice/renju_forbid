# Ren_forbid

[![ren_forbid compliant](https://img.shields.io/badge/Ren_fobid-realjustice-green.svg)](https://github.com/realjustice/renju_forbid)

最全五子棋禁手逻辑实现

目前已实现以下逻辑：

1. 三三禁手
2. 四四禁手
3. 长连禁手

## 内容列表

- [安装](#安装)
- [快速开始](#快速开始)
- [示例](#示例)
- [相关仓库](#相关仓库)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 安装

本项目使用 [go1.15](https://gomirrors.org/)。请确保你本地安装了它们。

```sh
$ go get -u github.com/realjustice/renju_forbid
```

## 快速开始

```sh
$ cat example/main.go
```

```go
package main

import (
	"fmt"
	"github.com/realjustice/renju_forbid"
)

func main() {
	// 0： 无禁手 1：三三禁手，2：四四禁手，3：长连
	forbidResult := ren.CheckForbid("(;SZ[15]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[dd];W[hd];B[de];W[ke];B[ci];W[li];B[df];W[lf];B[dg];W[mb];B[dh])")

	// 0：无结果，1：黑胜，-1：白胜
	OverResult := ren.CheckWin("(;SZ[15]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[dd];W[hd];B[de];W[ke];B[ci];W[li];B[df];W[lf];B[dg];W[mb];B[dh])")

	fmt.Printf("禁手结果：%d\n", forbidResult)
	fmt.Printf("对局结果：%d\n", OverResult)
}
```

```sh
# run example/main.go
$ go run example/main.go
```

## 示例

***可以调用下面的方法检查棋谱中是否包含禁手***

```go
result := ren.CheckForbid("your sgf")
```

会得到三种结果，**0： 无禁手 1：三三禁手，2：四四禁手，3：长连**

***可以调用下面的方法判断对局是否结束***

```go
result := ren.CheckWin("your sgf")
```

更多细节，请参考 [example](https://github.com/realjustice/renju_forbid/tree/master/example/)。

## 相关仓库

- [larry_dev/goban](https://gitee.com/larry_dev/goban?_from=gitee_search) —  最全围棋SGF解析器

## 维护者

[@realjustice](https://github.com/realjustice)。

## 如何贡献

[提一个 Issue](https://github.com/RichardLitt/standard-readme/issues/new) 或者提交一个 Pull Request，或者发送邮件至 [z_s_c_p@163.com](z_s_c_p@163.com)


## 使用许可

[MIT](LICENSE) © realjustice