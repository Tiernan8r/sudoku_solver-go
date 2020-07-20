package sudoku_solver

import (
	"errors"
	"fmt"
	"math"
)

type Grid struct {
	size  int
	board [][]*Cell
}

type Coordinate struct {
	RowIndex    int
	ColumnIndex int
	Value       int
}

func CreateGrid(size int) *Grid {
	boardCells := CreateBlankBoard(size)
	g := Grid{size, boardCells}
	return &g
}

func CreatePopulatedGrid(size int, coordinates []Coordinate) (*Grid, error) {
	blankGrid := CreateGrid(size)
	for _, coord := range coordinates {
		x, y := coord.RowIndex, coord.ColumnIndex
		if x > size {
			return nil, errors.New("row index integer out of bounds")
		} else if y > size {
			return nil, errors.New("column index integer out of bounds")
		}
		cell := blankGrid.board[x][y]
		cell.value = coord.Value
		cell.choices[coord.Value] = 2
		cell.solved = true
	}
	return blankGrid, nil
}

func (g *Grid) GetBoardSize() int {
	return g.size
}

func (g *Grid) GetBoard() [][]*Cell {
	return g.board
}

func (g *Grid) GetCell(rowIndex, columnIndex int) (*Cell, error) {
	if rowIndex > g.size {
		return nil, errors.New("row index integer out of bounds")
	} else if columnIndex > g.size {
		return nil, errors.New("column index integer out of bounds")
	}

	return g.board[rowIndex][columnIndex], nil
}

func (g *Grid) SetCell(rowIndex, columnIndex int, cell *Cell) error {

	if rowIndex > g.size {
		return errors.New("row index integer out of bounds")
	} else if columnIndex > g.size {
		return errors.New("column index integer out of bounds")
	}

	g.board[rowIndex][columnIndex] = cell
	return nil

}

func (g *Grid) SetNewCell(rowIndex, columnIndex int, val int, choices map[int]int, solved bool) error {
	cell := CreateCell(val, choices, solved)
	return g.SetCell(rowIndex, columnIndex, cell)
}

func (g *Grid) SetNewCellValue(rowIndex, columnIndex, val int) error {
	choices := make(map[int]int)
	return g.SetNewCell(rowIndex, columnIndex, val, choices, true)
}

func (g *Grid) GetRow(rowIndex int) ([]*Cell, error) {
	if rowIndex > g.size {
		return nil, errors.New("row index integer out of bounds")
	}

	row := make([]*Cell, 0)
	actualRow := g.board[rowIndex]
	for _, cell := range actualRow {
		row = append(row, cell)
	}
	return row, nil
}

func (g *Grid) GetColumn(columnIndex int) ([]*Cell, error) {

	if columnIndex > g.size {
		return nil, errors.New("column index integer out of bounds")
	}

	column := make([]*Cell, 0)

	for i := 0; i < g.size; i++ {
		cell := g.board[i][columnIndex]
		column = append(column, cell)
	}
	return column, nil
}

func (g *Grid) GetBox(rowIndex, columnIndex int) ([][]*Cell, error) {
	if rowIndex > g.size {
		return nil, errors.New("row index integer out of bounds")
	} else if columnIndex > g.size {
		return nil, errors.New("column index integer out of bounds")
	}

	boxSize := int(math.Sqrt(float64(g.size)))

	lowestRowIndex := int(math.Floor(float64(rowIndex / boxSize))) * boxSize
	lowestColumnIndex := int(math.Floor(float64(columnIndex / boxSize))) * boxSize
	//fmt.Println("ROW INDEX:", rowIndex)
	//fmt.Println("COLUMN INDEX:", columnIndex)
	//fmt.Println("BOX SIZE:", boxSize)
	//fmt.Println("LOWEST ROW INDEX:", lowestRowIndex)
	//fmt.Println("LOWEST COLUMN INDEX:", lowestColumnIndex)
	//fmt.Println()

	box := make([][]*Cell, boxSize)
	for i := 0; i < boxSize; i++ {
		box[i] = make([]*Cell, boxSize)
		for j := 0; j < boxSize; j++ {
			c := g.board[lowestRowIndex+i][lowestColumnIndex+j]
			box[i][j] = c
		}
	}
	return box, nil
}

func (g *Grid) RelativeCells(rowIndex, columnIndex int) ([]*Cell, error) {

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

	encounteredCell := make(map[*Cell]bool)
	var allCells []*Cell
	for _, rowCell := range row {
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

func (g *Grid) Display() {

	for _, row := range g.board {
		text := "|"
		for _, c := range row {
			val := " "
			if c.solved {
				val = fmt.Sprintf("%d", c.value)
			}
			text += fmt.Sprintf("%2s |", val)
		}
		for j := 0; j < g.size; j++ {
			fmt.Print("+---")
		}
		fmt.Println("+")
		fmt.Println(text)
	}
	for j := 0; j < g.size; j++ {
		fmt.Print("+---")
	}
	fmt.Println("+")

}
