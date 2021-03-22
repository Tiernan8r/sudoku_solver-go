package solver

import (
	"fmt"

	"github.com/tiernan8r/sudoku_solver/pkg/board"
)

// Solves for the unknown values in the given given Grid, and returns a pointer to the solved for Grid.
func Solve(board_ptr *board.Grid) *board.Grid {

	//Setup all the valid options for each cell
	board_ptr = CalculateChoices(board_ptr)

	// Count the total number of solved cells, and iterate until the number of solved cells matches the total number
	// of cells, once it does, the grid is solved
	numSolvedCells := 0
	// the board is a symmetric square, so the total number of cells is the square of the width
	totalNumCells := board_ptr.Size * board_ptr.Size
	solvedACell := false

	for numSolvedCells < totalNumCells {
		// reset the count of solved cells each iteration, to recount the total every loop
		numSolvedCells = 0
		// iterate over the cells in the grid
		for _, row := range board_ptr.Board {
			for _, cell := range row {

				// if the cell is solved, count it to the total number of solved
				if cell.Solved {
					numSolvedCells++
					continue
				}

				// count the number of possible values the cell could have
				numChoices := 0
				// The current value the cell is assumed to have, or could be assigned to have
				currentChoice := -1
				// iterate over all possible values for a cell
				for i := board_ptr.MinValue; i <= board_ptr.Size; i++ {
					// get the value assigned to the choices map for this potential value
					choiceValue := cell.Choices[i]
					// if the value from the map is 0 it means this entry is up for grabs, so count it.
					if choiceValue == 0 {
						// count the num of choices available, and set the current choice to be the iterated on one.
						numChoices++
						currentChoice = i
					}
				}
				// if there was only one available choice, set it.
				if numChoices == 1 {
					cell.SetValue(currentChoice)
					solvedACell = true
				}
			}
		}

		if !solvedACell {
			fmt.Println("Unable to solve a cell this iteration, entering an infinite loop, so quitting...")
			break
		}

		// recalculate the choice every total iteration
		board_ptr = CalculateChoices(board_ptr)
	}

	return board_ptr
}

// Calculates the potential call values in each cell based off of the solved for values in the row, column and box
// relative to each cell.
func CalculateChoices(board_ptr *board.Grid) *board.Grid {

	// iterate over all cells
	for rowIndex, row := range board_ptr.Board {
		for columnIndex, cell_ptr := range row {
			// get the set value for the given cell, even if it is not solved for
			cellValue := cell_ptr.Value
			// get all cells in the row, column, box relative to the current cell
			relativeCells, relativeCellsError := board_ptr.RelativeCells(rowIndex, columnIndex)
			// If it errored out, return the error.
			if relativeCellsError != nil {
				fmt.Println(relativeCellsError)
				return nil
			}
			// iterate over all the relative cells to set the potential values the cell could be
			for _, relativeCell_ptr := range relativeCells {
				// if the cell is solved for, we don't need to set the choices.
				if relativeCell_ptr.Solved {
					continue
				}
				// if the current cell is solved for, we reflect that in the choice of value for the relative cell
				if cell_ptr.Solved {
					// set the value in the choice map if it is unset.
					if relativeCell_ptr.Choices[cellValue] < 2 {
						// 1 signifies that the cell_ptr Value is taken by another cell in the row/column/box
						relativeCell_ptr.Choices[cellValue] = 1
					}
				}
			}
		}
	}

	////Display the choices for each cell in the grid
	//for rowIndex, row := range board_ptr.board {
	//	for columnIndex, cell_ptr := range row {
	//		fmt.Printf("(%d, %d): %d: %v\n", rowIndex, columnIndex, cell_ptr.choices, cell_ptr.solved)
	//	}
	//}

	return board_ptr
}
