package ren

import (
	"fmt"
	"gitee.com/larry_dev/goban"
	"regexp"
	"strconv"
	"strings"
)

// 棋盘大小
const S = 15

// 执黑
const BLACK_COLOR = 1

const WhITE_COLOR = -1

// 是否结束
func CheckWin(sgf string) int {
	return checkWin(initBoard(sgf))
}

// 校验是否含有禁手
func CheckForbid(sgf string) int {
	return checkForbid(initBoard(sgf))
}

// 坐标转为棋盘中的坐标字符串
func coordinateToRenjuPos(x, y, size int) string {
	if x == -1 && y == -1 {
		return "a0"
	}
	return strings.ToUpper(fmt.Sprintf("%s%d", string(x+'a'), size-y))
}

// 初始化棋盘
func initBoard(sgf string) ([S][S]int, int, int) {
	pos := convertSgfToPos(sgf)
	fmt.Println(pos)
	compile := regexp.MustCompile(`[a-o][0-9]+`)
	board := [S][S]int{}
	if compile == nil {
		return board, 0, 0
	}
	subMatch := compile.FindAllStringSubmatch(pos, -1)
	x, y := -1, -1
	curColor := BLACK_COLOR
	for _, v := range subMatch {
		curPos := v[0]
		x = int(curPos[0]) - 'a'
		pos, _ := strconv.Atoi(curPos[1:])
		if pos != 0 {
			y = pos - 1
			board[x][y] = curColor
		}
		curColor = -curColor
	}
	return board, x, y
}

func checkForbid(board [S][S]int, x, y int) int {
	copyBoard := board[:]
	line4V := newLine4V(copyBoard)

	if copyBoard[x][y] == 1 {
		return line4V.foulr(x, y)
	} else {
		return 0
	}
}

// 白胜的逻辑（五子连珠）
// 黑胜的逻辑（五子连珠且不包含禁手）
func checkWin(board [S][S]int, x, y int) int {
	copyBoard := board[:]
	line4V := newLine4V(copyBoard)

	// 黑落在禁手处,直接判白胜
	if checkForbid(board, x, y) != 0 {
		return WHITE_WIN
	}

	if copyBoard[x][y] == 1 && line4V.hasWon(x, y) {
		return BLACK_WIN
	} else if copyBoard[x][y] == -1 && line4V.hasWon(x, y) {
		return WHITE_WIN
	}

	return NO_RESULT
}

// Sgf转坐标
func convertSgfToPos(sgf string) string {
	kifu := goban.ParseSgf(sgf)
	pos := ""
	kifu.EachNode(func(n *goban.Node, move int) bool {
		if move == 0 {
			return false
		}
		pos += coordinateToRenjuPos(n.X, n.Y, S)
		return false
	})

	return strings.ToLower(pos)
}

func PrintBoard(board [S][S]int) {
	for _, v1 := range board {
		for k2, v2 := range v1 {
			if v2 == BLACK_COLOR {
				fmt.Print("X ")
			} else if v2 == WHITE_WIN {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
			if k2 == S-1 {
				fmt.Println()
			}
		}
	}
}
