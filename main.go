package main

import (
        "github.com/tiernan8r/sudoku_solver/solver";
        "github.com/tiernan8r/sudoku_solver/grid";
        "github.com/tiernan8r/sudoku_solver/cell";
        "fmt"
        )

func main() {
    fmt.Println("HEllo WORLD!")

    board_size := 9
    cells := [board_size][board_size]cell.Cell

    sudoku_board := grid.Grid{size: board_size, board: cells}

    sudoku_board = solver.Solve(sudoku_board)

    fmt.Println("FINITO!")

}
