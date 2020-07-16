package sudoku_solver

import (
        "math";
        )

type Grid struct {

    size int;
    board [size][size]cell.Cell;

}

func (grid *Grid) GetRow(row_index int) [grid.size]cell.Cell {
    // TODO: check for row index out of bounds.
    return grid.board[row_index];
}

func (grid *Grid) GetColumn(column_index int) [grid.size]cell.Cell {
    // TODO: check for column index out of bounds.
    column := make([]cell.Cell, grid.size)

    for i := 0; i < grid.size; i++ {
        c := grid.board[i][column_index]
        column = append(column, c)
    }

    return column
}

func (grid *Grid) GetBox(row_index, column_index int) [grid.size]cell.Cell {
    // TODO: check for index out of range.
    grid_size := grid.size
    box_size = math.Sqrt(grid_size)

    lowest_row_index := row_index % box_size
    lowest_column_index := column_index % box_size

    box := [box_size][box_size]cell.Cell
    for i := 0; i < box_size; i++ {
        for j := 0; j < box_size; j++ {
                c := grid.board[lowest_row_index + i][lowest_column_index + j]
                box[i][j] = c
        }
    }

    return box
}


