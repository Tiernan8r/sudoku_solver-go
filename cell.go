package sudoku_solver

type Cell struct {
	// the Value in this cell, default is 0, which means unset since 0 is never used in sudoku.
	value int
	// Choices is a map of "cell entry" -> "choice availability":
	// The "cell entry" are all the possible values that the cell could be (1-9) for a normal sudoku board.
	// "choice availabilty" is an integer:
	// - 0: the Value is up for grabs by this cell
	// - 1: the Value exists somewhere in the row/column/box relative to this cell.
	// - 2: the Value is solved/set as a hint
	choices map[int]int
	// whether this entry has been either preset or solved.
	solved bool
}

func (c *Cell) GetSolved() bool {
	return c.solved
}

func (c *Cell) SetSolved(solved bool) {
	c.solved = solved
}

func (c *Cell) GetChoices() *map[int]int {
	if c.choices == nil {
		c.choices = make(map[int]int)
	}
	return &c.choices
}

func (c *Cell) GetValue() int {
	return c.value
}

func (c *Cell) SetValue(val int) {
	c.value = val
	//choice := *c.GetChoices()
	//choice[val] = 2
	c.choices[val] = 2
	c.solved = true
	//c.SetSolved(true)
}

func CreateCell(val int, valueChoices map[int]int, solved bool) *Cell {
	// since a Value of 0 is the default Value, it isn't explicitly set.
	if val != 0 {
		valueChoices[val] = 2
	}
	c := Cell{value: val, choices: valueChoices, solved: solved}
	return &c
}

func CreateDefaultCell() *Cell {
	defaultChoices := make(map[int]int)
	return CreateCell(0, defaultChoices, false)
}

func CreateBlankBoard(boardSize int) [][]*Cell {

	cells := make([][]*Cell, boardSize)
	for i, _ := range cells {
		cells[i] = make([]*Cell, boardSize)
		for cellIndex, _ := range cells[i] {
			cells[i][cellIndex] = CreateDefaultCell()
		}
	}

	return cells
}
