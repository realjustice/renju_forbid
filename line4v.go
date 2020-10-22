package ren

type line4V struct {
	x1 [][S]int
	x2 [][S]int
	x3 [][S]int
	x4 [][S]int
}

const BLACK_WIN = 1
const NO_RESULT = 0
const WHITE_WIN = -1

const NO_BANNER_HANDS = 0
const DOUBLE3 = 1
const DOUBLE4 = 2
const OVERLINE = 3

func GA(x int) int {
	return x & 0xff
}

func GB(x int) int {
	return (x & 0xff00) >> 8
}

func newLine4V(board [][S]int) *line4V {
	line4 := &line4V{}
	line4.x1 = make([][S]int, S)
	line4.x2 = make([][S]int, S)
	line4.x3 = make([][S]int, S*2-1)
	line4.x4 = make([][S]int, S*2-1)
	if board == nil {
		return nil
	}
	var i, j int
	for i = 0; i < S; i++ {
		for j = 0; j < S; j++ {
			line4.x1[i][j] = board[i][j]
			line4.x2[i][j] = board[j][i]
		}
	}

	for i = 0; i < S; i++ {
		for j = 0; j <= i; j++ {
			line4.x3[i][j] = board[i-j][j]
			line4.x4[i][j] = board[i-j][S-1-j]
		}
		for ; j < S; j++ {
			line4.x3[i][j] = 1024
			line4.x4[i][j] = 1024
		}
	}
	for i = S; i < 2*S-1; i++ {
		for j = 0; j <= i-S; j++ {
			line4.x3[i][j] = 1024
			line4.x4[i][j] = 1024
		}
		for j = i - S + 1; j < S; j++ {
			line4.x3[i][j] = board[i-j][j]
			line4.x4[i][j] = board[i-j][S-1-j]
		}
	}

	return line4
}

func (line4 *line4V) hasWon(x, y int) bool {
	sign := 0
	if line4.A5(x, y) != 0 {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return true
	}
	return false
}

func (line4 *line4V) foulr(x, y int) int {
	sign := 0
	if line4.x1[x][y] == 0 {
		sign = 1
		line4.x1[x][y] = 1
		line4.x2[y][x] = 1
		line4.x3[x+y][y] = 1
		line4.x4[S-1-y+x][S-1-y] = 1
	}

	if line4.x1[x][y] == WhITE_COLOR {
		return NO_BANNER_HANDS
	}

	if line4.A5(x, y) != NO_BANNER_HANDS {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return NO_BANNER_HANDS
	}

	if line4.double4(x, y) != NO_BANNER_HANDS {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return DOUBLE4
	}

	if line4.double3r(x, y) != NO_BANNER_HANDS {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return DOUBLE3
	}

	if line4.overline(x, y) != NO_BANNER_HANDS {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return OVERLINE
	}

	line4.x1[x][y] -= sign
	line4.x2[y][x] -= sign
	line4.x3[x+y][y] -= sign
	line4.x4[S-1-y+x][S-1-y] -= sign

	return NO_BANNER_HANDS

}

func (line4 *line4V) A5(x, y int) int {
	if line4.x1[x][y] == 0 {
		return 0
	}
	l1 := newLine(line4.x1[x][:])
	l2 := newLine(line4.x2[y][:])
	l3 := newLine(line4.x3[x+y][:])
	l4 := newLine(line4.x4[S-1-y+x][:])
	p1, p2, p3, p4 := y, x, y, S-1-y
	if l1.A5(p1) != 0 {
		return 1
	}
	if l2.A5(p2) != 0 {
		return 1
	}
	if l3.A5(p3) != 0 {
		return 1
	}
	if l4.A5(p4) != 0 {
		return 1
	}
	return 0

}

func (line4 *line4V) double4(x int, y int) int {
	if line4.B4(x, y) >= 2 {
		return 1
	}
	return 0
}

func (line4 *line4V) B4(x, y int) int {
	l1 := newLine(line4.x1[x][:])
	l2 := newLine(line4.x2[y][:])
	l3 := newLine(line4.x3[x+y][:])
	l4 := newLine(line4.x4[S-1-y+x][:])
	p1, p2, p3, p4 := y, x, y, S-1-y
	count := 0
	count += l1.B4(p1)
	count += l2.B4(p2)
	count += l3.B4(p3)
	count += l4.B4(p4)
	return count
}

func (line4 *line4V) double3r(x, y int) int {
	if line4.A3r(x, y) >= 2 {
		return 1
	}
	return 0
}

func (line4 *line4V) A3r(x, y int) int {
	l1 := newLine(line4.x1[x][:])
	l2 := newLine(line4.x2[y][:])
	l3 := newLine(line4.x3[x+y][:])
	l4 := newLine(line4.x4[S-1-y+x][:])
	p1, p2, p3, p4 := y, x, y, S-1-y
	count := 0

	ll1 := l1.A3(p1)
	ll2 := l2.A3(p2)
	ll3 := l3.A3(p3)
	ll4 := l4.A3(p4)

	if ll1 != 0 {
		if ll1 < 65536 {

			r := GA(ll1)
			//		cout << r << '\n' ;
			if line4.x1[x][y] == 1 && (line4.foulr(x, r) != 0 || line4.A5test(x, r) != 0) {
				count--
			}
			count++
		} else {

			r1 := GA(ll1)
			r2 := GB(ll1)
			if line4.x1[x][y] == 1 && (line4.foulr(x, r1) != 0 || line4.A5test(x, r1) != 0) && (line4.foulr(x, r2) != 0 || line4.A5test(x, r2) != 0) {
				count--
			}
			count++
		}
	}
	if ll2 != 0 {
		if ll2 < 65536 {

			r := GA(ll2)
			//		cout << r << '\n' ;
			if line4.x1[x][y] == 1 && (line4.foulr(r, y) != 0 || line4.A5test(r, y) != 0) {
				count--
			}
			count++
		} else {
			r1 := GA(ll2)
			r2 := GB(ll2)
			if line4.x1[x][y] == 1 && (line4.foulr(r1, y) != 0 || line4.A5test(r1, y) != 0) && (line4.foulr(r2, y) != 0 || line4.A5test(r2, y) != 0) {
				count--
			}
			count++
		}
	}
	if ll3 != 0 {
		if ll3 < 65536 {
			r := GA(ll3)
			//		cout << r << '\n' ;
			if line4.x1[x][y] == 1 && (line4.foulr(x+y-r, r) != 0 || line4.A5test(x+y-r, r) != 0) {
				count--
			}
			count++
		} else {
			r1 := GA(ll3)
			r2 := GB(ll3)
			if line4.x1[x][y] == 1 && (line4.foulr(x+y-r1, r1) != 0 || line4.A5test(x+y-r1, r1) != 0) && (line4.foulr(x+y-r2, r2) != 0 || line4.A5test(x+y-r2, r2) != 0) {
				count--
			}
			count++
		}
	}
	if ll4 != 0 {
		if ll4 < 65536 {
			r := GA(ll4)
			if line4.x1[x][y] == 1 && (line4.foulr(S-1+x-y-r, S-1-r) != 0 || line4.A5test(S-1+x-y-r, S-1-r) != 0) {
				count--
			}
			count++
		} else {
			r1 := GA(ll4)
			r2 := GB(ll4)
			if line4.x1[x][y] == 1 && (line4.foulr(S-1+x-y-r1, S-1-r1) != 0 || line4.A5test(S-1+x-y-r1, S-1-r1) != 0) && (line4.foulr(S-1+x-y-r2, S-1-r2) != 0 || line4.A5test(S-1+x-y-r2, S-1-r2) != 0) {
				count--
			}
			count++
		}
	}

	return count
}

func (line4 *line4V) A5test(x, y int) int {
	sign := 0
	if line4.x1[x][y] == 0 {
		sign = 1
		line4.x1[x][y] = 1
		line4.x2[y][x] = 1
		line4.x3[x+y][y] = 1
		line4.x4[S-1-y+x][S-1-y] = 1
	}
	l1 := newLine(line4.x1[x][:])
	l2 := newLine(line4.x2[y][:])
	l3 := newLine(line4.x3[x+y][:])
	l4 := newLine(line4.x4[S-1-y+x][:])
	p1, p2, p3, p4 := y, x, y, S-1-y
	if l1.A5(p1) != 0 {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return 1
	}
	if l2.A5(p2) != 0 {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return 1
	}
	if l3.A5(p3) != 0 {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return 1
	}
	if l4.A5(p4) != 0 {
		line4.x1[x][y] -= sign
		line4.x2[y][x] -= sign
		line4.x3[x+y][y] -= sign
		line4.x4[S-1-y+x][S-1-y] -= sign
		return 1
	}
	line4.x1[x][y] -= sign
	line4.x2[y][x] -= sign
	line4.x3[x+y][y] -= sign
	line4.x4[S-1-y+x][S-1-y] -= sign
	return 0
}

func (line4 *line4V) overline(x, y int) int {
	if line4.A6(x, y) != 0 {
		return 1
	}
	return 0
}

func (line4 *line4V) A6(x, y int) int {
	l1 := newLine(line4.x1[x][:])
	l2 := newLine(line4.x2[y][:])
	l3 := newLine(line4.x3[x+y][:])
	l4 := newLine(line4.x4[S-1-y+x][:])
	p1, p2, p3, p4 := y, x, y, S-1-y
	if l1.A6(p1) != 0 {
		return 1
	}
	if l2.A6(p2) != 0 {
		return 1
	}
	if l3.A6(p3) != 0 {
		return 1
	}
	if l4.A6(p4) != 0 {
		return 1
	}
	return 0
}
