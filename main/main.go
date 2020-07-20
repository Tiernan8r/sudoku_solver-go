package main

import (
	"fmt"
	"github.com/tiernan8r/sudoku_solver"
)

func main() {
	const boardSize int = 9
	//const boardSize int = 16

	sudokuBoard_ptr := sudoku_solver.CreateGrid(boardSize)

	coordinates := make([]sudoku_solver.Coordinate, 0)
	// 9 * 9 GRID
	coordinates = append(coordinates, sudoku_solver.Coordinate{0, 0, 6})
	coordinates = append(coordinates, sudoku_solver.Coordinate{0, 1, 4})
	coordinates = append(coordinates, sudoku_solver.Coordinate{0, 4, 3})
	coordinates = append(coordinates, sudoku_solver.Coordinate{0, 8, 7})

	coordinates = append(coordinates, sudoku_solver.Coordinate{1, 0, 5})
	coordinates = append(coordinates, sudoku_solver.Coordinate{1, 2, 1})
	coordinates = append(coordinates, sudoku_solver.Coordinate{1, 4, 7})
	coordinates = append(coordinates, sudoku_solver.Coordinate{1, 6, 9})

	coordinates = append(coordinates, sudoku_solver.Coordinate{2, 7, 1})

	coordinates = append(coordinates, sudoku_solver.Coordinate{3, 2, 4})
	coordinates = append(coordinates, sudoku_solver.Coordinate{3, 3, 9})
	coordinates = append(coordinates, sudoku_solver.Coordinate{3, 5, 8})
	coordinates = append(coordinates, sudoku_solver.Coordinate{3, 7, 6})

	coordinates = append(coordinates, sudoku_solver.Coordinate{4, 1, 8})
	coordinates = append(coordinates, sudoku_solver.Coordinate{4, 5, 3})
	coordinates = append(coordinates, sudoku_solver.Coordinate{4, 7, 2})

	coordinates = append(coordinates, sudoku_solver.Coordinate{5, 3, 4})

	coordinates = append(coordinates, sudoku_solver.Coordinate{6, 0, 4})
	coordinates = append(coordinates, sudoku_solver.Coordinate{6, 3, 1})
	coordinates = append(coordinates, sudoku_solver.Coordinate{6, 4, 5})
	coordinates = append(coordinates, sudoku_solver.Coordinate{6, 5, 7})
	coordinates = append(coordinates, sudoku_solver.Coordinate{6, 7, 3})

	coordinates = append(coordinates, sudoku_solver.Coordinate{7, 0, 2})
	coordinates = append(coordinates, sudoku_solver.Coordinate{7, 2, 8})
	coordinates = append(coordinates, sudoku_solver.Coordinate{7, 3, 3})
	coordinates = append(coordinates, sudoku_solver.Coordinate{7, 7, 4})

	coordinates = append(coordinates, sudoku_solver.Coordinate{8, 0, 7})
	coordinates = append(coordinates, sudoku_solver.Coordinate{8, 1, 5})
	coordinates = append(coordinates, sudoku_solver.Coordinate{8, 7, 9})
	coordinates = append(coordinates, sudoku_solver.Coordinate{8, 8, 6})

	// 16 * 16 GRID
	// 1-9, A=10,B=11,C=12,D=13,E=14,F=15
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,0,5})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,2,14})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,4,10})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,7,15})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,9,3})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,11,11})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{0,12,4})
	//
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,0,15})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,1,3})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,2,13})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,3,9})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,5,4})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,8,7})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,14,1})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{1,15,5})
	//
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,0,6})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,4,0})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,5,0})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,6,0})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,8,0})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,9,0})
	//coordinates = append(coordinates, sudoku_solver.Coordinate{2,9,0})

	sudokuBoard_ptr, _ = sudoku_solver.CreatePopulatedGrid(boardSize, coordinates)

	fmt.Println("INPUT:")
	sudokuBoard_ptr.Display()
	fmt.Println("============")
	sudokuBoard_ptr = sudoku_solver.Solve(sudokuBoard_ptr)

	fmt.Println("ANSWER:")
	sudokuBoard_ptr.Display()
	fmt.Println("============")

}
