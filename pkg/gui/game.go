package gui

import (
	"bytes"
	"fmt"

	"github.com/ZongBen/GoFive/pkg/game"
)

var board = ""
var Turn = ""
var buffer = new(bytes.Buffer)

func RenderBoard(b game.Board) {
	board = ""
	buffer.Reset()
	if b.GetTurn() {
		Turn = "Black"
	} else {
		Turn = "White"
	}
	board += "Welcome to GoFive!\n"
	board += "Use 'w', 'a', 's', 'd' to move the cursor and 'e' to place a piece.\n"
	board += "Press 'q' to quit.\n"
	board += "Turn: " + Turn + "\n"
	for y := 0; y < b.GetHeight(); y++ {
		renderLine(b, y)
	}
	buffer.WriteString(board)
	fmt.Print(buffer.String())
}

func renderLine(b game.Board, row int) {
	for part := 0; part < 3; part++ {
		for col := 0; col < b.GetWidth(); col++ {
			renderSwitcher(b, col, row, part)
		}
		board += "\n"
	}
}

func renderSwitcher(b game.Board, x, y, part int) {
	piece := b.GetPoint(x, y)
	select_x, select_y := b.GetSelectorPosition()
	isSelected := x == select_x && y == select_y
	switch piece.State {
	case game.EMPTY:
		renderPosition(part, isSelected)
		break
	case game.BLACK:
		renderBlack(part, isSelected)
		break
	case game.WHITE:
		renderWhite(part, isSelected)
		break
	}
}

func renderPosition(part int, isSelected bool) {
	if part == 0 {
		if isSelected {
			board += "┏     ┓"
		} else {
			board += "       "
		}
	} else if part == 1 {
		if isSelected {
			board += "┃  +  ┃"
		} else {
			board += "   +   "
		}
	} else {
		if isSelected {
			board += "┗     ┛"
		} else {
			board += "       "
		}
	}
}

func renderBlack(part int, isSelected bool) {
	if part == 0 {
		if isSelected {
			board += "┏ *** ┓"
		} else {
			board += "  ***  "
		}
	} else if part == 1 {
		if isSelected {
			board += "┃*****┃"
		} else {
			board += " ***** "
		}
	} else {
		if isSelected {
			board += "┗ *** ┛"
		} else {
			board += "  ***  "
		}
	}
}

func renderWhite(part int, isSelected bool) {
	if part == 0 {
		if isSelected {
			board += "┏ OOO ┓"
		} else {
			board += "  OOO  "
		}
	} else if part == 1 {
		if isSelected {
			board += "┃OOOOO┃"
		} else {
			board += " OOOOO "
		}
	} else {
		if isSelected {
			board += "┗ OOO ┛"
		} else {
			board += "  OOO  "
		}
	}
}