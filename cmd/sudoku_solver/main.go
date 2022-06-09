package main

import (
	"fmt"

	"github.com/tiernan8r/sudoku_solver/pkg/board"
	"github.com/tiernan8r/sudoku_solver/pkg/solver"
)

const boardSize = 9

func main() {
	sudokuBoard_ptr := board.CreateGrid(boardSize)

	coordinates := make([]board.Coordinate, 0)
	// 9 * 9 GRID
	coordinates = append(coordinates, board.Coordinate{0, 1, 5})
	coordinates = append(coordinates, board.Coordinate{0, 3, 7})
	coordinates = append(coordinates, board.Coordinate{0, 4, 2})
	coordinates = append(coordinates, board.Coordinate{0, 6, 9})
	coordinates = append(coordinates, board.Coordinate{0, 8, 3})

	coordinates = append(coordinates, board.Coordinate{1, 0, 9})
	coordinates = append(coordinates, board.Coordinate{1, 5, 6})
	coordinates = append(coordinates, board.Coordinate{1, 8, 4})

	coordinates = append(coordinates, board.Coordinate{2, 4, 1})

	coordinates = append(coordinates, board.Coordinate{3, 1, 2})
	coordinates = append(coordinates, board.Coordinate{3, 3, 4})
	coordinates = append(coordinates, board.Coordinate{3, 4, 6})
	coordinates = append(coordinates, board.Coordinate{3, 8, 8})

	coordinates = append(coordinates, board.Coordinate{4, 1, 7})
	coordinates = append(coordinates, board.Coordinate{4, 2, 4})
	coordinates = append(coordinates, board.Coordinate{4, 3, 8})

	coordinates = append(coordinates, board.Coordinate{5, 5, 1})
	coordinates = append(coordinates, board.Coordinate{5, 7, 7})

	coordinates = append(coordinates, board.Coordinate{6, 1, 9})
	coordinates = append(coordinates, board.Coordinate{6, 8, 5})

	coordinates = append(coordinates, board.Coordinate{7, 0, 5})
	coordinates = append(coordinates, board.Coordinate{7, 2, 1})
	coordinates = append(coordinates, board.Coordinate{7, 7, 2})
	coordinates = append(coordinates, board.Coordinate{7, 8, 6})

	// 16 * 16 GRID
	// 1-9, A=10,B=11,C=12,D=13,E=14,F=15
	//coordinates = append(coordinates, board.Coordinate{0,0,5})
	//coordinates = append(coordinates, board.Coordinate{0,2,14})
	//coordinates = append(coordinates, board.Coordinate{0,4,10})
	//coordinates = append(coordinates, board.Coordinate{0,7,15})
	//coordinates = append(coordinates, board.Coordinate{0,9,3})
	//coordinates = append(coordinates, board.Coordinate{0,11,11})
	//coordinates = append(coordinates, board.Coordinate{0,12,4})
	//
	//coordinates = append(coordinates, board.Coordinate{1,0,15})
	//coordinates = append(coordinates, board.Coordinate{1,1,3})
	//coordinates = append(coordinates, board.Coordinate{1,2,13})
	//coordinates = append(coordinates, board.Coordinate{1,3,9})
	//coordinates = append(coordinates, board.Coordinate{1,5,4})
	//coordinates = append(coordinates, board.Coordinate{1,8,7})
	//coordinates = append(coordinates, board.Coordinate{1,14,1})
	//coordinates = append(coordinates, board.Coordinate{1,15,5})
	//
	//coordinates = append(coordinates, board.Coordinate{2,0,6})
	//coordinates = append(coordinates, board.Coordinate{2,4,0})
	//coordinates = append(coordinates, board.Coordinate{2,5,0})
	//coordinates = append(coordinates, board.Coordinate{2,6,0})
	//coordinates = append(coordinates, board.Coordinate{2,8,0})
	//coordinates = append(coordinates, board.Coordinate{2,9,0})
	//coordinates = append(coordinates, board.Coordinate{2,9,0})

	sudokuBoard_ptr, _ = board.CreatePopulatedGrid(boardSize, coordinates)

	fmt.Println("INPUT:")
	sudokuBoard_ptr.Display()
	fmt.Println("============")
	sudokuBoard_ptr = solver.Solve(sudokuBoard_ptr)

	fmt.Println("ANSWER:")
	sudokuBoard_ptr.Display()
	fmt.Println("============")

}
