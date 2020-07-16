package sudoku

import (
        "math";
        )

type Grid struct {

    size int;
    board [][]Cell;

}

func (g *Grid) GetRow(row_index int) []Cell {
    // TODO: check for row index out of bounds.
    return g.board[row_index];
}

func (g *Grid) GetColumn(column_index int) []Cell {
    // TODO: check for column index out of bounds.
    column := make([]Cell, g.size)

    for i := 0; i < g.size; i++ {
        c := g.board[i][column_index]
        column = append(column, c)
    }

    return column
}

func (g *Grid) GetBox(row_index, column_index int) [][]Cell {
    // TODO: check for index out of range.
    box_size := int(math.Sqrt(float64(g.size)))

    lowest_row_index := row_index % box_size
    lowest_column_index := column_index % box_size

    box := make([][]Cell, box_size)
    for i := 0; i < box_size; i++ {
        box[i] = make([]Cell, box_size)
        for j := 0; j < box_size; j++ {
                c := g.board[lowest_row_index + i][lowest_column_index + j]
                box[i][j] = c
        }
    }

    return box
}


