package main

import (
	"fmt"
	"github.com/realjustice/renju_forbid"
)

func main() {
	// demo 1 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[dh];W[ch];B[eh];W[an];B[fh];W[dn];B[hg];W[go];B[hi];W[ln];B[ih];W[nn];B[kh];W[la];B[hh]) 否
	// demo 2 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[hg];W[nn];B[hi];W[cn];B[gh];W[jn];B[ih];W[ho];B[lh];W[en];B[dh];W[hn];B[hh]) 否
	// demo 3 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[eh];W[eo];B[fh];W[lo];B[hg];W[hn];B[hi];W[nn];B[ih];W[no];B[jh];W[kc];B[hh]) 否
	// demo 4 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[fe];W[ed];B[bo];W[aa];B[ff];W[nn];B[fh];W[oo];B[fi];W[lo];B[fj];W[ej];B[gg];W[hn];B[hg];W[fa];B[if];W[am];B[ih];W[je];B[na];W[ji];B[ig]) 是
	// demo 5 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[eh];W[en];B[fh];W[bm];B[ig];W[go];B[ii];W[lo];B[jh];W[jo];B[lh];W[no];B[ih]) 是
	// demo 6 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[eh];W[cn];B[hh];W[go];B[ig];W[lo];B[ii];W[cb];B[kh];W[bl];B[mh];W[fn];B[ih]) 否
	// demo 7 略（换个方向）
	// demo 8 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[dh];W[co];B[gh];W[gn];B[mn];W[kn];B[hg];W[mc];B[hi];W[fb];B[jh];W[bl];B[mh];W[em];B[hh]) 是
	// demo 9 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[eh];W[dn];B[fh];W[mn];B[hg];W[hn];B[hi];W[da];B[ih];W[ka];B[jh];W[kh];B[hh]) 否
	// demo 10 (;GM[1]FF[4]CA[UTF-8]SZ[15]AP[弈客围棋];B[fe];W[ed];B[ff];W[an];B[fh];W[en];B[fi];W[nb];B[fj];W[ej];B[mn];W[ia];B[gg];W[nd];B[hg];W[ah];B[if];W[ho];B[ih];W[je];B[lg];W[ji];B[ig]) 否

	// 0： 无禁手 1：三三禁手，2：四四禁手，3：长连
	forbidResult := ren.CheckForbid("(;SZ[15]KM[7.5]HA[0]AW[fj][gi][gh][gf][hg][ig][jh][lg][le]AB[lh][ij][hk][gj][hh][gg][hf][if][ki][kh];W[ji];B[kj];W[kg];B[jg];W[he];B[hj];W[jj])")

	// 0：无结果，1：黑胜，-1：白胜
	OverResult := ren.CheckWin("(;SZ[15]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[dd];B[ed];B[fd];B[ec];W[];B[ee])")

	// 0: 无结果，1：黑五子连珠，-1：白：五子连珠
	renjuResult := ren.IsRenju("(;AP[WGo.js:2]FF[4]GM[1]CA[UTF-8]SZ[15];B[dd];W[];B[fd];W[];B[ec];W[];B[ee];W[];B[ed];W[];B[ef];W[];B[eg])")

	fmt.Printf("禁手结果：%d\n", forbidResult)
	fmt.Printf("对局结果：%d\n", OverResult)
	fmt.Printf("五子连珠结果：%d\n", renjuResult)
}
