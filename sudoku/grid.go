package sudoku

import "math"

// TODO: change this implementation
const board_size int = 9
//const box_size int = 3

type Grid struct {
	size  int
	board [][]Cell
}

func (g *Grid) GetRow(row_index int) []Cell {
	// TODO: check for row index out of bounds.
	row := make([]Cell, board_size)
	actual_row := g.board[row_index]
	for _, c := range actual_row {
		row = append(row, c)
	}
	return row
}

func (g *Grid) GetColumn(column_index int) []Cell {
	// TODO: check for column index out of bounds.
	//column := [board_size]Cell{}
	column := make([]Cell, board_size)

	for i := 0; i < g.size; i++ {
		c := g.board[i][column_index]
		column = append(column, c)
	}

	return column
}

func (g *Grid) GetBox(row_index, column_index int) [][]Cell {
	// TODO: check for index out of range.
	box_size := int(math.Sqrt(float64(board_size)))

	lowest_row_index := row_index % box_size
	lowest_column_index := column_index % box_size

	box := make([][]Cell, box_size)
	for i := 0; i < box_size; i++ {
		box[i] = make([]Cell, box_size)
		for j := 0; j < box_size; j++ {
			c := g.board[lowest_row_index+i][lowest_column_index+j]
			box[i][j] = c
		}
	}

	return box
}
