package sudoku_solver

import "fmt"

func Solve(populatedBoard_ptr *Grid) *Grid {

	//Setup all the valid options for each cell
	populatedBoard_ptr = CalculateChoices(populatedBoard_ptr)

	// Find the cell with only one option, eliminate the choice from others in row/column/box
	populatedBoard_ptr = eliminatePossibilities(populatedBoard_ptr)

	//return
	return populatedBoard_ptr
}

func CalculateChoices(board_ptr *Grid) *Grid {

	// Setup all the valid options for each cell
	for rowIndex, row := range board_ptr.board {
		for columnIndex, cell_ptr := range row {
			cellValue := cell_ptr.value
			allCells, allCellsError := board_ptr.RelativeCells(rowIndex, columnIndex)
			if allCellsError != nil {
				fmt.Println(allCellsError)
				return nil
			}
			for _, currentCell_ptr := range allCells {
				if currentCell_ptr.solved {
					continue
				}
				if cell_ptr.solved {
					if currentCell_ptr.choices[cellValue] < 2 {
						// 1 signifies that the cell_ptr Value is taken by another cell in the row/column/box
						currentCell_ptr.choices[cellValue] = 1
					}
				}
			}
		}
	}

	//for rowIndex, row := range board_ptr.board {
	//	for columnIndex, cell_ptr := range row {
	//		fmt.Printf("(%d, %d): %d: %v\n", rowIndex, columnIndex, cell_ptr.choices, cell_ptr.solved)
	//	}
	//}

	return board_ptr

}

func eliminatePossibilities(board_ptr *Grid) *Grid {

	numSolvedCells := 0
	totalNumCells := board_ptr.size * board_ptr.size

	for numSolvedCells < totalNumCells {
		numSolvedCells = 0
		for rowIndex, row := range board_ptr.board {
			for columnIndex, currentCell := range row {

				if currentCell.solved {
					numSolvedCells++
				}

				allCells, _ := board_ptr.RelativeCells(rowIndex, columnIndex)
				for _, cell := range allCells {
					if cell.solved || cell.value != 0 {
						continue
					}
					numAvailableValues := 0
					currentChoice := 0
					for i := 1; i <= board_ptr.size; i++ {
						choiceValue := cell.choices[i]
						if choiceValue == 0 {
							numAvailableValues++
							currentChoice = i
						}
					}
					if numAvailableValues == 1 {
						cell.SetValue(currentChoice)
					}
				}
			}
		}
		board_ptr = CalculateChoices(board_ptr)
	}

	return board_ptr
}
