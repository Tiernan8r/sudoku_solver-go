package main

import (
    "fmt"
    "github.com/tiernan8r/sudoku_solver/sudoku"
)

func main() {
    fmt.Println("HEllo WORLD!")

    const board_size int = 9
    cells := [board_size][board_size]sudoku.Cell{}

    sudoku_board := sudoku.Grid{size: board_size, board: cells}

    sudoku_board = sudoku.Solve(sudoku_board)

    fmt.Println("FINITO!")

}
