/*
   Backtracing N-queen problem
*/

package main

import (
	"fmt"
)

func printSolution(b [][]int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Println("")
	}
}

// check if this is a solution for the queenPosition
func checkQSolution(b [][]int, queenRow int, queenCol int) bool {
	// check if there is another queen on the same row
	for c := 0; c < len(b); c++ {
		if c != queenCol && b[queenRow][c] == 1 {
			return false
		}
	}

	for r := 0; r < len(b); r++ {
		if r != queenRow && b[r][queenCol] == 1 {
			return false
		}
	}

	// first diag
	cleft, cright := queenCol, queenCol
	rup, rdown := queenRow, queenRow
	for {
		m1, m2 := 0, 0
		if cleft > 0 && rup < len(b)-1 {
			cleft--
			rup++
			m1++
		}

		if cright < len(b)-1 && rdown > 0 {
			cright++
			rdown--
			m2++
		}

		if m1 == 0 && m2 == 0 {
			break
		}

		if b[rup][cleft] == 1 && m1 > 0 {
			return false
		}

		if b[rdown][cright] == 1 && m2 > 0 {
			return false
		}
	}

	// second diag
	cleft, cright = queenCol, queenCol
	rup, rdown = queenRow, queenRow
	for {
		m1, m2 := 0, 0
		if cleft > 0 && rdown > 0 {
			// going down & left
			cleft--
			rdown--
			m1++
		}

		if cright < len(b)-1 && rup < len(b)-1 {
			// going up & right
			cright++
			rup++
			m2++
		}

		if m1 == 0 && m2 == 0 {
			break
		}

		if b[rdown][cleft] == 1 && m1 > 0 {
			return false
		}

		if b[rup][cright] == 1 && m2 > 0 {
			return false
		}
	}

	return true
}

func solve(b [][]int, solutions *int, col int) bool {
	if col >= len(b) {
		(*solutions)++
	}

	for r := 0; r < len(b); r++ {
		if checkQSolution(b, r, col) {
			b[r][col] = 1

			if solve(b, solutions, col+1) {
				return true
			}

			b[r][col] = 0
		}
	}

	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	n = 10

	board := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		r := make([]int, 0, n)
		for j := 0; j < n; j++ {
			r = append(r, 0)
		}
		board = append(board, r)
	}

	solutions := 0
	solve(board, &solutions, 0)
	fmt.Printf("solutions %d\n", solutions)
}
