package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a 2 dimensional array
type Matrix struct {
	rows [][]int
}

// Rows returns the count of rows in a matrix
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.rows))
	for i := range rows {
		rows[i] = make([]int, len(m.rows[0]))
	}

	for i := 0; i < len(m.rows); i++ {
		for j := 0; j < len(m.rows[i]); j++ {
			rows[i][j] = m.rows[i][j]
		}
	}
	return rows
}

// Cols returns the count of rows in a matrix
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len(m.rows[0]))
	for i := range cols {
		cols[i] = make([]int, len(m.rows))
	}

	for i := 0; i < len(m.rows); i++ {
		for j := 0; j < len(m.rows[i]); j++ {
			cols[j][i] = m.rows[i][j]
		}
	}
	return cols
}

// Set sets value of cell and returns success
func (m *Matrix) Set(r int, c int, val int) bool {
	if r >= 0 && r < len(m.rows) {
		if c >= 0 && c < len(m.rows[0]) {
			m.rows[r][c] = val
			return true
		}
	}
	return false
}

// New returns a new matrix struct
func New(in string) (*Matrix, error) {
	var m Matrix
	ccount := -1
	for _, row := range strings.Split(strings.Trim(in, " "), "\n") {
		c := strings.Split(strings.Trim(row, " "), " ")
		if ccount == -1 {
			ccount = len(c)
		}
		if ccount != len(c) {
			return &m, errors.New("Column lengths are irregular")
		}

		ci := make([]int, len(c))
		for i := 0; i < len(c); i++ {
			val, err := strconv.Atoi(c[i])
			ci[i] = val
			if err != nil {
				return &m, err
			}
		}
		m.rows = append(m.rows, ci)
	}
	return &m, nil
}
