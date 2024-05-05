package field

import (
	"fmt"
	"strings"
)

type Field struct {
	state *[3][3]string
}

func NewField() *Field {
	return &Field{
		state: new([3][3]string),
	}
}

func (f *Field) String() string {

	builder := strings.Builder{}

	for _, row := range f.state {
		for _, cell := range row {

			if cell == "" {
				cell = " "
			}

			builder.WriteString("[")
			builder.WriteString(cell)
			builder.WriteString("] ")

		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (f *Field) Put(sign string, x, y int) error {

	if err := checkCoords(x, y); err != nil {
		return err
	}

	if !f.Empty(x, y) {
		return fmt.Errorf("place [%d,%d] is already reserved", x, y)
	}

	f.state[x-1][y-1] = sign

	return nil
}

func (f *Field) Empty(x, y int) bool {

	if !validCellAddress(x) || !validCellAddress(y) {
		return false
	}

	return f.state[x-1][y-1] == ""
}

func (f *Field) Sign(x, y int) (string, error) {

	if err := checkCoords(x, y); err != nil {
		return "", err
	}

	return f.state[x-1][y-1], nil
}

func checkCoords(x, y int) error {

	if !validCellAddress(x) || !validCellAddress(y) {
		return fmt.Errorf("invalid cell address [%d,%d]", x, y)
	} else {
		return nil
	}
}

func validCellAddress(i int) bool {
	return i >= 1 && i <= 3
}
