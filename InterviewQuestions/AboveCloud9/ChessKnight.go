package main

import "math"

var board = [8][8]int{}
var minSteps int = math.MaxInt

func is_valid(startX int, startY int) bool {
	if (startX >= 0 && startX < 8) && (startY >= 0 && startY < 8) {
		return true
	} else {
		return false
	}
}

func solve(startX int, startY int, endX int, endY int, board [][]int, step int) {
	if startX == endX && startY == endY {
		minSteps = min(minSteps, step)
		return 
	} else if !is_valid(startX, startY) || board[startX][startY] != 0 {
		return
	} else {
		board[startX][startY] = 1
		solve(startX+3, startY+2, endX, endY, board, step+1)
		solve(startX+3, startY-2, endX, endY, board, step+1)
		solve(startX-3, startY+2, endX, endY, board, step+1)
		solve(startX-3, startY-2, endX, endY, board, step+1)
		solve(startX+2, startY+3, endX, endY, board, step+1)
		solve(startX+2, startY-3, endX, endY, board, step+1)
		solve(startX-2, startY+3, endX, endY, board, step+1)
		solve(startX-2, startY-3, endX, endY, board, step+1)
		board[startX][startY] = 0
	}
}
