package board

import (
	"errors"
	"fmt"
	"math"
)

type Grid struct {
	MinValue int
	Size     int
	Board    [][]*Cell
}

type Coordinate struct {
	RowIndex    int
	ColumnIndex int
	Value       int
}

// Creates a blank Grid object containing blank Cells
func CreateGrid(size int) *Grid {
	boardCells := CreateBlankBoard(size)
	minValue := 0
	if size == 9 {
		minValue = 1
	}
	g := Grid{minValue, size, boardCells}
	return &g
}

// Create a Grid object pre populated with cell values in the given positions
func CreatePopulatedGrid(size int, coordinates []Coordinate) (*Grid, error) {
	// Get a blank grid to begin with
	blankGrid := CreateGrid(size)
	// Iterate over the given coordinates
	for _, coord := range coordinates {
		x, y := coord.RowIndex, coord.ColumnIndex
		if x > size {
			return nil, errors.New("row index integer out of bounds")
		} else if y > size {
			return nil, errors.New("column index integer out of bounds")
		}

		// Get the cells associated with the given coordinates and assign the values accordingly.
		cell := blankGrid.Board[x][y]
		cell.Value = coord.Value
		// A value of 2 in the map means the cell is solved for this value
		cell.Choices[coord.Value] = 2
		cell.Solved = true
	}
	return blankGrid, nil
}

func (g *Grid) GetCell(rowIndex, columnIndex int) (*Cell, error) {
	if rowIndex > g.Size {
		return nil, errors.New("row index integer out of bounds")
	} else if columnIndex > g.Size {
		return nil, errors.New("column index integer out of bounds")
	}

	return g.Board[rowIndex][columnIndex], nil
}

func (g *Grid) SetCell(rowIndex, columnIndex int, cell *Cell) error {

	if rowIndex > g.Size {
		return errors.New("row index integer out of bounds")
	} else if columnIndex > g.Size {
		return errors.New("column index integer out of bounds")
	}

	g.Board[rowIndex][columnIndex] = cell
	return nil

}

func (g *Grid) SetNewCell(rowIndex, columnIndex int, val int, choices map[int]int, solved bool) error {
	cell := CreateCell(val, choices, solved)
	return g.SetCell(rowIndex, columnIndex, cell)
}

// A method to set the value of the cell at the given indices to the given value
func (g *Grid) SetNewCellValue(rowIndex, columnIndex, val int) error {
	choices := make(map[int]int)
	return g.SetNewCell(rowIndex, columnIndex, val, choices, true)
}

func (g *Grid) GetRow(rowIndex int) ([]*Cell, error) {
	if rowIndex > g.Size {
		return nil, errors.New("row index integer out of bounds")
	}

	row := make([]*Cell, 0)
	actualRow := g.Board[rowIndex]
	for _, cell := range actualRow {
		row = append(row, cell)
	}
	return row, nil
}

func (g *Grid) GetColumn(columnIndex int) ([]*Cell, error) {

	if columnIndex > g.Size {
		return nil, errors.New("column index integer out of bounds")
	}

	column := make([]*Cell, 0)

	// Iterate down the rows to get the cell in the given column.
	for i := 0; i < g.Size; i++ {
		cell := g.Board[i][columnIndex]
		column = append(column, cell)
	}
	return column, nil
}

// In sudoku, each cell occupies a box where the entry has to be unique, this method finds the box for the cell.
func (g *Grid) GetBox(rowIndex, columnIndex int) ([][]*Cell, error) {
	if rowIndex > g.Size {
		return nil, errors.New("row index integer out of bounds")
	} else if columnIndex > g.Size {
		return nil, errors.New("column index integer out of bounds")
	}

	// Find the width of the box from the grid size.
	boxSize := int(math.Sqrt(float64(g.Size)))

	// Find the upper left corner of the box to index from
	lowestRowIndex := int(math.Floor(float64(rowIndex/boxSize))) * boxSize
	lowestColumnIndex := int(math.Floor(float64(columnIndex/boxSize))) * boxSize

	box := make([][]*Cell, boxSize)
	for i := 0; i < boxSize; i++ {
		box[i] = make([]*Cell, boxSize)
		for j := 0; j < boxSize; j++ {
			c := g.Board[lowestRowIndex+i][lowestColumnIndex+j]
			box[i][j] = c
		}
	}
	return box, nil
}

// A method to get all the unique cells in the row, column and box that the given cell indices occupy.
func (g *Grid) RelativeCells(rowIndex, columnIndex int) ([]*Cell, error) {

	// get the row, column and box of the cell
	row, rowError := g.GetRow(rowIndex)
	column, columnError := g.GetColumn(columnIndex)
	box, boxError := g.GetBox(rowIndex, columnIndex)

	if rowError != nil {
		return nil, rowError
	} else if columnError != nil {
		return nil, columnError
	} else if boxError != nil {
		return nil, boxError
	}

	// keep track of whether the cell is unique or not to prevent double counting.
	encounteredCell := make(map[*Cell]bool)
	var allCells []*Cell
	// iterate over the cells in the row, column and box and add the cells to the total only if they are unique.
	for _, rowCell := range row {
		// the uniqueness check is technically superfluous for the row since the slice is initially empty...
		if encounteredCell[rowCell] == false {
			encounteredCell[rowCell] = true
			allCells = append(allCells, rowCell)
		}
	}
	for _, columnCell := range column {
		if encounteredCell[columnCell] == false {
			encounteredCell[columnCell] = true
			allCells = append(allCells, columnCell)
		}
	}
	for _, boxRow := range box {
		for _, boxCell := range boxRow {
			if encounteredCell[boxCell] == false {
				encounteredCell[boxCell] = true
				allCells = append(allCells, boxCell)
			}
		}
	}

	return allCells, nil
}

// Creates a simple text table display of the sudoku grid.
func (g *Grid) Display() {

	// Iterate down the rows
	for _, row := range g.Board {
		// setup the border of the table with a pipe
		text := "|"
		// iterate across the cells in the row
		for _, cell := range row {
			// If the cell is unsolved, display only a space character
			val := " "
			if cell.Solved {
				// if the cell is solved for, convert the int value to a string
				val = fmt.Sprintf("%d", cell.Value)
			}
			// display each value in a box with a pipe separating each value
			text += fmt.Sprintf("%2s |", val)
		}
		// Print the header for the row cells
		for j := 0; j < g.Size; j++ {
			fmt.Print("+---")
		}
		fmt.Println("+")
		// print the cell values.
		fmt.Println(text)
	}
	// print a base for the cells.
	for j := 0; j < g.Size; j++ {
		fmt.Print("+---")
	}
	fmt.Println("+")

}
