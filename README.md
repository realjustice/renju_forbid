# Ren_forbid

最全五子棋禁手逻辑实现，使用 go 语言编写

***可以调用下面的方法检查棋谱中是否包含禁手***

```go
result := ren.CheckForbid("(;GM[1]FF[4]CA[UTF-8]SZ[15];B[fe];W[ed];B[bo];W[aa];B[ff];W[nn];B[fh];W[oo];B[fi];W[lo];B[fj];W[ej];B[gg];W[hn];B[hg];W[fa];B[if];W[am];B[ih];W[je];B[na];W[ji];B[ig])")// 棋谱sgf文件
```

会得到三种结果，**0： 无禁手 1：三三禁手，2：四四禁手，3：长连**



***可以调用下面的方法判断对局是否结束***

```go
result := ren.CheckWin("(;GM[1]FF[4]CA[UTF-8]SZ[15];B[fe];W[ed];B[bo];W[aa];B[ff];W[nn];B[fh];W[oo];B[fi];W[lo];B[fj];W[ej];B[gg];W[hn];B[hg];W[fa];B[if];W[am];B[ih];W[je];B[na];W[ji];B[ig])")// 棋谱sgf文件
```

会得到三种结果，**0：无结果，1：黑胜，-1：白胜**

更多细节，请参考example