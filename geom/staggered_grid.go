package geom

type SGrid[TCell any, TEdge any] struct {
	Width, Height int
	Cells         [][]TCell
	Edges         [][][2]TEdge
}

func NewSGrid[TCell any, TEdge any](width, height int, defaultCell func() TCell, defaultEdge func() TEdge) *SGrid[TCell, TEdge] {
	g := SGrid[TCell, TEdge]{
		Width:  width,
		Height: height,
		Cells:  make([][]TCell, width),
		Edges:  make([][][2]TEdge, width+1),
	}
	for x := range g.Cells {
		g.Cells[x] = make([]TCell, height)
		g.Edges[x] = make([][2]TEdge, height+1)
		for y := 0; y < height; y++ {
			g.Cells[x][y] = defaultCell()
			g.Edges[x][y] = [2]TEdge{defaultEdge(), defaultEdge()}
		}
		g.Edges[x][height] = [2]TEdge{defaultEdge(), defaultEdge()}
	}
	g.Edges[width] = make([][2]TEdge, height+1)
	for y := 0; y < height; y++ {
		g.Edges[width][y] = [2]TEdge{defaultEdge(), defaultEdge()}
	}
	g.Edges[width][height] = [2]TEdge{defaultEdge(), defaultEdge()}
	return &g
}

func (g *SGrid[TCell, TEdge]) SetCell(x, y int, cell TCell) {
	g.Cells[x][y] = cell
}

func (g *SGrid[TCell, TEdge]) GetCell(x, y int) TCell {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return *new(TCell)
	}
	return g.Cells[x][y]
}

func (g *SGrid[TCell, TEdge]) GetNeighbors(x, y int) [4]TCell {
	pt := Point2{x, y}
	return [4]TCell{
		g.GetCell(pt.Offset(N).XY()),
		g.GetCell(pt.Offset(E).XY()),
		g.GetCell(pt.Offset(S).XY()),
		g.GetCell(pt.Offset(W).XY()),
	}
}

func (g *SGrid[TCell, TEdge]) SetEdges(x, y int, n, e, s, w TEdge) {
	g.Edges[x][y][0] = w
	g.Edges[x][y][1] = s
	g.Edges[x][y+1][1] = n
	g.Edges[x+1][y][0] = e
}

func (g *SGrid[TCell, TEdge]) GetEdges(x, y int) (TEdge, TEdge, TEdge, TEdge) {
	return g.Edges[x][y+1][1], g.Edges[x+1][y][0], g.Edges[x][y][1], g.Edges[x][y][0]
}

func (g *SGrid[TCell, TEdge]) Swap(x1, y1, x2, y2 int) {
	g.Cells[x1][y1], g.Cells[x2][y2] = g.GetCell(x2, y2), g.GetCell(x1, y1)
}
