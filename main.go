package main

import (
        "github.com/tiernan8r/sudoku_solver/sudoku";
        "fmt"
        )

func main() {
    fmt.Println("HEllo WORLD!")

    board_size := 9
    cells := [board_size][board_size]Cell

    sudoku_board := Grid{size: board_size, board: cells}

    sudoku_board = Solve(sudoku_board)

    fmt.Println("FINITO!")

}
