package text

import (
	"fmt"
	"slices"

	"github.com/relvox/iridescence_go/geom"
)

var TAB_SIZE = 4

type TextGrid struct {
	Lines  []string
	Cursor geom.Point2
}

func NewTextGrid() *TextGrid {
	return &TextGrid{
		Lines:  []string{""},
		Cursor: [2]int{},
	}
}

func (g *TextGrid) AppendRune(r rune) {
	// fmt.Println(g.Cursor, string(r))
	switch c := r; {
	case c == 8: // backspace
		if g.Cursor.X() == 0 {
			if g.Cursor.Y() == 0 {
				break
			}
			g.Cursor[0] = len(g.Lines[g.Cursor.Y()-1])
			g.Lines[g.Cursor.Y()-1] += g.Lines[g.Cursor.Y()]
			g.Cursor[1]--
			if len(g.Lines)-1 > g.Cursor.Y()+1 {
				copy(g.Lines[g.Cursor.Y()+1:], g.Lines[g.Cursor.Y()+2:])
			}
			g.Lines = g.Lines[:len(g.Lines)-1]
			break
		}
		p := g.Cursor.X()
		if p == len(g.Lines[g.Cursor.Y()]) {
			g.Lines[g.Cursor.Y()] = g.Lines[g.Cursor.Y()][:p-1]
		} else {
			g.Lines[g.Cursor.Y()] = g.Lines[g.Cursor.Y()][:p] + g.Lines[g.Cursor.Y()][p+1:]
		}
		g.Cursor[0]--
	case c == '\t':
		for i := 0; i < TAB_SIZE; i++ {
			g.AppendRune(' ')
		}
	case c == '\n':
		g.Lines = append(g.Lines, "")
		g.Cursor[0] = 0
		g.Cursor[1] += 1
	// case c >= ' ' && c <= '~':
	case c == 56:
	case c == '\r':
		panic(fmt.Errorf("unsupported rune: \\r [%d]", r))
	case c == 127: // delete
		panic(fmt.Errorf("unsupported rune: \\DEL [%d]", r))
	default:
		g.Lines[g.Cursor.Y()] += string(c)
		g.Cursor[0] += 1
		// log.Println(fmt.Errorf("unknown rune: %c [%d]", r, r))
	}
}

func (g *TextGrid) AppendString(s string) {
	for _, r := range s {
		g.AppendRune(r)
	}
}

func (g *TextGrid) Print(a ...any) {
	g.AppendString(fmt.Sprint(a...))
}

func (g *TextGrid) Printf(format string, a ...any) {
	g.AppendString(fmt.Sprintf(format, a...))
}

func (g *TextGrid) Println(a ...any) {
	g.AppendString(fmt.Sprintln(a...))
}

func (g *TextGrid) RewindTo(cur geom.Point2) {
	newX, newY := cur.XY()
	oldX, oldY := g.Cursor.XY()
	if oldY < newY {
		return
	} else if oldY == newY {
		if oldX <= newX {
			return
		} else {
			g.Lines[g.Cursor.Y()] = g.Lines[g.Cursor.Y()][:newX+1]
			g.Cursor[0] = newX
		}
	} else {
		g.Lines = g.Lines[:newY+1]
		g.Cursor[1] = newY
		g.Lines[g.Cursor.Y()] = g.Lines[g.Cursor.Y()][:newX]
		g.Cursor[0] = newX
	}
}

// BoxedLines
// width, height, font - in units
// marginSpaceWrap - in runes
func (g *TextGrid) BoxedLines(width, height, font, marginSpaceWrap int, flipBreak bool) []string {
	rWidth, rHeight := width/font-2, height/font/2-1
	var results [][]string = make([][]string, 0, len(g.Lines))
	for _, line := range g.Lines {
		var subLines []string
		for len(line) > rWidth {
			var b int
			for b = rWidth; b > rWidth-marginSpaceWrap; b-- {
				if line[b] == ' ' {
					break
				}
			}
			if line[b] != ' ' {
				subLines, line = append(subLines, line[:rWidth]), line[rWidth:]
			} else {
				subLines, line = append(subLines, line[:b]), line[b+1:]
			}
		}
		if len(subLines) == 0 {
			results = append(results, []string{line})
			continue
		}
		if len(line) != 0 {
			subLines = append(subLines, line)
		}
		results = append(results, subLines)
	}

	var result []string
	for _, subLines := range results {
		if flipBreak {
			slices.Reverse(subLines)
		}
		result = append(result, subLines...)
	}
	h := len(result)
	if h > rHeight {
		return result[h-rHeight:]
	}
	return result
}
