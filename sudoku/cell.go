package sudoku

type Cell struct {
    value int
    // TODO: change this 9 to reflect grid size.
    choices []int
    solved bool
}

func CreateCell(val int, edittable bool, options ...int) *Cell {
        c := Cell{value: val, choices: options, solved: edittable}
        return &c
}

