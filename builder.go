package packetdiagram

import (
	"fmt"
	"io"
	"log"

	svg "github.com/ajstarks/svgo"
)

func Draw(def *Definition, out io.Writer) error {
	canvas := svg.New(out)
	dim := calculateDimensions(def)
	canvas.Start(dim.Canvas.Width, dim.Canvas.Height)
	defineStyles(def, dim, canvas)

	drawBackground(def, dim, canvas)
	drawXAxis(def, dim, canvas)
	drawYAxis(def, dim, canvas)
	drawPlacements(def, dim, canvas)
	canvas.End()
	return nil
}

type Dimensions struct {
	Canvas Dimension
	XAxis  Dimension
	YAxis  Dimension
	Cell   Dimension
}

type Dimension struct {
	Width  int
	Height int
}

func calculateDimensions(def *Definition) Dimensions {
	return Dimensions{
		Canvas: Dimension{
			Width:  1100,
			Height: 500,
		},
		XAxis: Dimension{
			Width:  1000,
			Height: 30,
		},
		YAxis: Dimension{
			Width:  50,
			Height: 500,
		},
		Cell: Dimension{
			Width:  30,
			Height: 30,
		},
	}
}

func defineStyles(def *Definition, dim Dimensions, canvas *svg.SVG) {
	defineStyleForXAxisBits(def, dim, canvas)
	defineStyleForPlacements(def, dim, canvas)
	defineStyleForBreakMark(def, dim, canvas)
}

func defineStyleForXAxisBits(def *Definition, dim Dimensions, canvas *svg.SVG) {
	canvas.Style("text/css", fmt.Sprintf(`
text.bit{
	fill:%s;
	font-size:%s;
	text-anchor:middle;
}`,
		def.GetTextColor(),
		def.GetTextSize(),
	),
	)
}

func defineStyleForPlacements(def *Definition, dim Dimensions, canvas *svg.SVG) {
	canvas.Style("text/css", fmt.Sprintf(`
polygon.placement{
	fill:white;
	stroke:black;
}

text.placement{
	fill:%s;
	font-family:%s;
	font-size:%s;
	text-anchor:middle;
}`,
		def.GetTextColor(),
		def.GetTextFontFamily(),
		def.GetTextSize(),
	),
	)
}

func defineStyleForBreakMark(def *Definition, dim Dimensions, canvas *svg.SVG) {
	canvas.Style("text/css", fmt.Sprintf(`
path.breakmark{
	fill:none;
	stroke:black;
}`,
		def.GetTextColor(),
		def.GetTextFontFamily(),
		def.GetTextSize(),
	),
	)
}

func drawBackground(def *Definition, dim Dimensions, canvas *svg.SVG) {
	canvas.Rect(0, 0, int(dim.Canvas.Width), int(dim.Canvas.Height), "id='background'", fmt.Sprintf("fill='%s'", def.GetBackgroundColor()), "stroke='none'")
}

func drawXAxis(def *Definition, dim Dimensions, canvas *svg.SVG) {
	if def.ShouldShowXAxisBits() {
		drawXAxisBits(def, dim, canvas)
	}
	if def.ShouldShowXAxisOctets() {
		drawXAxisOctets(def, dim, canvas)
	}
}

func drawXAxisBits(def *Definition, dim Dimensions, canvas *svg.SVG) {
	xs, ys, labels := calculateXAxisBitLabelDimensions(def, dim)

	for i := range xs {
		canvas.Text(xs[i], ys[i], labels[i], `class="bit"`)
	}
	canvas.Text(xs[len(xs)-1]+dim.Cell.Width, ys[len(ys)-1], "bit", `class="bit"`)
}

func calculateXAxisBitLabelDimensions(def *Definition, dim Dimensions) (xs []int, ys []int, labels []string) {
	count := int(def.OctetsPerLine * 8)

	xs = make([]int, count)
	ys = make([]int, count)
	labels = make([]string, count)

	startX := int(dim.YAxis.Width)
	for i := 0; i < count; i++ {
		xs[i] = startX + (i * int(dim.Cell.Width)) + (dim.Cell.Width / 2)
		ys[i] = int(dim.Cell.Height) - 4
		labels[i] = fmt.Sprintf("%d", i+int(def.GetXAxisBitsOrigin()))
	}
	return
}

func drawXAxisOctets(def *Definition, dim Dimensions, canvas *svg.SVG) {

}

func drawYAxis(def *Definition, dim Dimensions, canvas *svg.SVG) {

}

type Cursor struct {
	x int
	y int
}

func drawPlacements(def *Definition, dim Dimensions, canvas *svg.SVG) {
	cur := &Cursor{x: 0, y: 0}

	for i, p := range def.Placements {
		log.Printf("placement: %#v", p)
		drawPlacement(def, dim, cur, p, i, canvas)
	}
}

func drawPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement, index int, canvas *svg.SVG) {
	polygons := getPlacementPolygons(def, dim, cur, p)
	if len(polygons) == 0 {
		return
	}

	for _, polygon := range polygons {
		canvas.Polygon(polygon.xs, polygon.ys, `class="placement"`)
		if p.VariableLength {
			drawBreakMark(def, polygon, canvas)
		}
		drawPlacementText(def, dim, p, polygon, canvas)
	}
}

func drawBreakMark(def *Definition, polygon Polygon, canvas *svg.SVG) {
	left, top, right, bottom := polygon.findBoundingBox()

	sx, sy, cx, cy, px, py, ex, ey := getBreakMarkPoints(left, (top+bottom)/2-3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(left, (top+bottom)/2+3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)

	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(right, (top+bottom)/2-3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(right, (top+bottom)/2+3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
}

func getBreakMarkPoints(x, y int) (sx, sy, cx, cy, px, py, ex, ey int) {
	sx = x - 10
	cx = x
	px = x
	ex = x + 10

	sy = y
	cy = y - 10
	py = y + 10
	ey = y
	return
}

func drawPlacementText(def *Definition, dim Dimensions, p Placement, polygon Polygon, canvas *svg.SVG) {
	left, top, right, bottom := polygon.findBoundingBox()
	canvas.Text((left+right)/2, ((top+bottom)/2)+(dim.Cell.Height/6), p.Label, `class="placement"`)
}

type Polygon struct {
	xs []int
	ys []int
}

func (p Polygon) findBoundingBox() (left int, top int, right int, bottom int) {
	left = p.xs[0]
	top = p.ys[0]
	right = p.xs[0]
	bottom = p.ys[0]
	for i := range p.xs {
		if p.xs[i] < left {
			left = p.xs[i]
		}
		if right < p.xs[i] {
			right = p.xs[i]
		}
		if p.ys[i] < top {
			top = p.ys[i]
		}
		if bottom < p.ys[i] {
			bottom = p.ys[i]
		}
	}
	return
}

func getPlacementPolygons(def *Definition, dim Dimensions, cur *Cursor, p Placement) []Polygon {
	bitsPerLine := def.GetBitsPerLine()
	bitDistributions := getBitDistributions(bitsPerLine, cur, p)
	defer advanceCursor(def, bitDistributions, cur)

	lb := len(bitDistributions)
	if lb == 1 {
		polygon := getSingleRectPlacement(def, dim, cur, p)
		return []Polygon{polygon}
	}

	if lb == 2 && int(p.Bits) <= def.GetBitsPerLine() {
		return getSeparatedTwoLinesPlacement(def, dim, cur, p, bitDistributions)
	}

	polygon := getMultipleLinesPlacement(def, dim, cur, p, bitDistributions)
	return []Polygon{polygon}
}

func advanceCursor(def *Definition, bitDistributions []int, cur *Cursor) {
	bitsPerLine := def.GetBitsPerLine()
	for _, bd := range bitDistributions {
		cur.x += bd
		if cur.x == bitsPerLine {
			cur.x = 0
			cur.y += 1
		}
	}
}

func getBitDistributions(bitsPerLine int, cur *Cursor, p Placement) []int {
	bitsToGo := int(p.Bits)
	availableInCurrentLine := bitsPerLine - cur.x

	if availableInCurrentLine >= bitsToGo {
		return []int{int(bitsToGo)}
	}

	result := make([]int, 0)
	result = append(result, availableInCurrentLine)
	bitsToGo -= availableInCurrentLine

	for bitsToGo > 0 {
		if bitsToGo < bitsPerLine {
			result = append(result, bitsToGo)
			break
		}
		result = append(result, bitsPerLine)
		bitsToGo -= bitsPerLine
	}

	return result
}

func doesFitInCurrentLine(def *Definition, cur *Cursor, p Placement) bool {
	return availableBitsInCurrentLine(def, cur) >= int(p.Bits)
}

func availableBitsInCurrentLine(def *Definition, cur *Cursor) int {
	return def.GetBitsPerLine() - cur.x
}

func getSingleRectPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement) Polygon {
	left := dim.YAxis.Width + (cur.x * dim.Cell.Width)
	right := left + (int(p.Bits) * dim.Cell.Width)
	top := dim.XAxis.Height + (cur.y * dim.Cell.Height)
	bottom := top + dim.Cell.Height
	return createRectPolygon(left, top, right, bottom)
}

func createRectPolygon(left, top, right, bottom int) Polygon {
	xs := make([]int, 5)
	xs[0] = left
	xs[1] = right
	xs[2] = right
	xs[3] = left
	xs[4] = left

	ys := make([]int, 5)
	ys[0] = top
	ys[1] = top
	ys[2] = bottom
	ys[3] = bottom
	ys[4] = top

	return Polygon{xs: xs, ys: ys}
}

func doesFitInTwoLines(def *Definition, cur *Cursor, p Placement) bool {
	return int(p.Bits)-availableBitsInCurrentLine(def, cur) < def.GetBitsPerLine()
}

func getSeparatedTwoLinesPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement, bitDistributions []int) []Polygon {
	cw := dim.Cell.Width
	ch := dim.Cell.Height
	/*
		          ┌──────────┐
		          │    (1)   │
		┌───────┐ └──────────┘
		│  (2)  │
		└───────┘
	*/
	left1 := dim.YAxis.Width + (cur.x * cw)
	right1 := left1 + bitDistributions[0]*cw
	top1 := dim.XAxis.Height + (cur.y * ch)
	bottom1 := top1 + ch
	p1 := createRectPolygon(left1, top1, right1, bottom1)

	left2 := dim.YAxis.Width
	right2 := left2 + bitDistributions[1]*cw
	top2 := dim.XAxis.Height + ((cur.y + 1) * ch)
	bottom2 := top2 + ch
	p2 := createRectPolygon(left2, top2, right2, bottom2)

	return []Polygon{p1, p2}
}

func getMultipleLinesPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement, bitDistributions []int) Polygon {
	cw := dim.Cell.Width
	ch := dim.Cell.Height
	/*
		pattern 1:
		      ┌───────────┐
		      │           │
		┌─────┘           │
		│                 │
		│                 │
		│                 │
		└─────────────────┘
	*/
	if bitDistributions[1] == def.GetBitsPerLine() {
		startX := dim.YAxis.Width + (cur.x * cw)
		startY := dim.XAxis.Height + (cur.y * ch)
		xs := make([]int, 7)
		xs[0] = startX
		xs[1] = xs[0] + (bitDistributions[0] * cw)
		xs[2] = xs[1]
		xs[3] = dim.YAxis.Width
		xs[4] = xs[3]
		xs[5] = xs[0]
		xs[6] = xs[0]

		ys := make([]int, 7)
		ys[0] = startY
		ys[1] = ys[0]
		ys[2] = ys[0] + (ch * len(bitDistributions))
		ys[3] = ys[2]
		ys[4] = ys[0] + ch
		ys[5] = ys[4]
		ys[6] = ys[0]

		return Polygon{xs: xs, ys: ys}
	}

	/*
		pattern 2:
		      ┌───────────┐
		      │           │
		┌─────┘           │
		│                 │
		│                 │
		│           ┌─────┘
		│           │
		└───────────┘
	*/
	startX := dim.YAxis.Width + (cur.x * cw)
	startY := dim.XAxis.Height + (cur.y * ch)
	xs := make([]int, 9)
	xs[0] = startX
	xs[1] = xs[0] + (bitDistributions[0] * cw)
	xs[2] = xs[1]
	xs[3] = dim.YAxis.Width + (bitDistributions[1] * cw)
	xs[4] = xs[3]
	xs[5] = dim.YAxis.Width
	xs[6] = xs[5]
	xs[7] = xs[0]
	xs[8] = xs[0]

	ys := make([]int, 9)
	ys[0] = startY
	ys[1] = ys[0]
	ys[2] = ys[0] + ((len(bitDistributions) - 1) * ch)
	ys[3] = ys[2]
	ys[4] = ys[3] + ch
	ys[5] = ys[4]
	ys[6] = ys[0] + ch
	ys[7] = ys[6]
	ys[8] = ys[0]

	return Polygon{xs: xs, ys: ys}
}
