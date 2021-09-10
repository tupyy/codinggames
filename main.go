/*
   Backtracing N-queen problem
*/

package main

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

func main() {
	// var n int
	// fmt.Scan(&n)

	// r := make([]int, 0, 4)
	// board := [][]int{r, r, r, r}
}
