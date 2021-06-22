package packetdiagram

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

func Draw(def *Definition, out io.Writer) error {
	canvas := svg.New(out)
	dim := calculateDimensions(def)
	canvas.Start(int(dim.Canvas.Width), int(dim.Canvas.Height))
	defineStyles(def, dim, canvas)

	drawBackground(def, dim, canvas)
	drawXAxis(def, dim, canvas)
	drawYAxis(def, dim, canvas)
	drawPlacements(def, dim, canvas)
	canvas.End()
	return nil
}

func drawBackground(def *Definition, dim Dimensions, canvas *svg.SVG) {
	canvas.Rect(0, 0, int(dim.Canvas.Width), int(dim.Canvas.Height), "id='background'", fmt.Sprintf("fill='%s'", def.GetBackgroundColor()), "stroke='none'")
}

func drawXAxis(def *Definition, dim Dimensions, canvas *svg.SVG) {
	if def.ShouldShowXAxisOctets() {
		drawXAxisOctets(def, dim, canvas)
	}
	if def.ShouldShowXAxisBits() {
		drawXAxisBits(def, dim, canvas)
	}
}

func drawXAxisBits(def *Definition, dim Dimensions, canvas *svg.SVG) {
	xs, ys, labels := calculateXAxisBitLabelDimensions(def, dim)
	h := int(def.GetXAxisBitsHeight())
	cw := int(dim.Cell.Width)
	lox := int(cw / 2)
	loy := int(def.GetXAxisBitsHeight() * 3 / 4)
	for i := range xs {
		canvas.Line(xs[i], ys[i], xs[i], ys[i]+h, `class="x-bit"`)
		canvas.Text(xs[i]+lox, ys[i]+loy, labels[i], `class="x-bit"`)
	}

	last := len(xs) - 1
	xlast := xs[last]
	ylast := ys[last]
	canvas.Line(xlast+cw, ylast, xlast+cw, ylast+h, `class="x-bit"`)
	canvas.Text(xs[0]+5, ys[0]+int(def.GetAxisTitleTextSizeInPixels()), "bit", `class="x-bit-title"`)
}

func calculateXAxisBitLabelDimensions(def *Definition, dim Dimensions) (xs []int, ys []int, labels []string) {
	count := int(def.GetOctetsPerLine() * 8)

	h := int(def.GetXAxisBitsHeight())
	cw := int(dim.Cell.Width)
	o := int(def.GetXAxisBitsOrigin())
	u := int(def.GetXAxisBitsUnit())
	startX := int(dim.YAxis.Width)

	xs = make([]int, count)
	ys = make([]int, count)
	labels = make([]string, count)

	for i := 0; i < count; i++ {
		xs[i] = startX + (i * cw)
		ys[i] = int(dim.XAxis.Height) - h
		if def.GetXAxisBitsDirection() == XAxisBitsDirectionLeftToRight {
			labels[i] = fmt.Sprintf("%d", o+(i%u))
		} else {
			labels[i] = fmt.Sprintf("%d", o+(u-(i%u)))
		}
	}
	return
}

func drawXAxisOctets(def *Definition, dim Dimensions, canvas *svg.SVG) {
	xs, ys, labels := calculateXAxisOctetLabelDimensions(def, dim)
	h := int(def.GetXAxisBitsHeight())
	cw := int(dim.Cell.Width)
	lox := int(cw * 4)                           // label offset x
	loy := int(def.GetXAxisBitsHeight()) * 3 / 4 // label offset y
	for i := range xs {
		canvas.Line(xs[i], ys[i], xs[i], ys[i]+h, `class="x-octet"`)
		canvas.Text(xs[i]+lox, ys[i]+loy, labels[i], `class="x-octet"`)
	}

	last := len(xs) - 1
	xlast := xs[last]
	ylast := ys[last]
	canvas.Line(xlast+(cw*8), ylast, xlast+(cw*8), ylast+h, `class="x-octet"`)
	canvas.Text(xs[0]+5, ys[0]+int(def.GetAxisTitleTextSizeInPixels())+3, "octet", `class="x-octet-title"`)
}

func calculateXAxisOctetLabelDimensions(def *Definition, dim Dimensions) (xs []int, ys []int, labels []string) {
	count := int(def.GetOctetsPerLine())

	cw := int(dim.Cell.Width)

	xs = make([]int, count)
	ys = make([]int, count)
	labels = make([]string, count)

	startX := int(dim.YAxis.Width)
	for i := 0; i < count; i++ {
		xs[i] = startX + (i * cw * 8)
		ys[i] = 0
		labels[i] = fmt.Sprintf("%d", i)
	}
	return
}

func drawYAxis(def *Definition, dim Dimensions, canvas *svg.SVG) {
	if def.ShouldShowYAxisOctets() {
		drawYAxisOctets(def, dim, canvas)
	}
	if def.ShouldShowYAxisBits() {
		drawYAxisBits(def, dim, canvas)
	}
}

func drawYAxisBits(def *Definition, dim Dimensions, canvas *svg.SVG) {
	xs, ys, labels := calculateYAxisBitLabelDimensions(def, dim)
	w := int(def.GetYAxisBitsWidth())

	lox := int(def.GetYAxisBitsWidth()) - 5
	loy := int(dim.Cell.Height * 3 / 4)

	canvas.Text(xs[0]+lox, ys[0]+int(def.GetAxisTitleTextSizeInPixels()+3), "bit", `class="y-bit-title"`)

	for i := range xs {
		canvas.Line(xs[i], ys[i], xs[i]+w, ys[i], `class="y-bit"`)
		canvas.Text(xs[i]+lox, ys[i]+loy, labels[i], `class="y-bit"`)
	}

	last := len(xs) - 1
	ch := int(def.GetCellHeight())
	canvas.Line(xs[last], ys[last]+ch, xs[last]+w, ys[last]+ch, `class="y-bit"`)
}

func calculateYAxisBitLabelDimensions(def *Definition, dim Dimensions) (xs, ys []int, labels []string) {
	offsetX := int(def.GetYAxisOctetsWidth())
	offsetY := int(dim.XAxis.Height)
	xs = make([]int, 0)
	ys = make([]int, 0)
	labels = make([]string, 0)
	totalBits := int(def.GetTotalPlacementBits())

	for curr := int(def.GetYAxisOctetsOrigin()); curr < totalBits; curr += int(def.GetBitsPerLine()) {
		xs = append(xs, int(offsetX))
		ys = append(ys, int(offsetY+int(dim.Cell.Height)*(curr/int(def.GetBitsPerLine()))))
		labels = append(labels, fmt.Sprintf("%d", curr))
	}
	return
}

func drawYAxisOctets(def *Definition, dim Dimensions, canvas *svg.SVG) {
	xs, ys, labels := calculateYAxisOctetLabelDimensions(def, dim)

	w := int(def.GetYAxisOctetsWidth())
	lox := int(def.GetYAxisOctetsWidth()) - 5
	loy := int(dim.Cell.Height * 3 / 4)

	canvas.Text(xs[0]+lox, ys[0]+int(def.GetAxisTitleTextSizeInPixels()+3), "octet", `class="y-octet-title"`)

	for i := range xs {
		canvas.Line(xs[i], ys[i], xs[i]+w, ys[i], `class="y-bit"`)
		canvas.Text(xs[i]+lox, ys[i]+loy, labels[i], `class="y-octet"`)
	}

	last := len(xs) - 1
	ch := int(dim.Cell.Height)
	canvas.Line(xs[last], ys[last]+ch, xs[last]+w, ys[last]+ch, `class="y-bit"`)
}

func calculateYAxisOctetLabelDimensions(def *Definition, dim Dimensions) (xs, ys []int, labels []string) {
	offsetX := 0
	offsetY := int(dim.XAxis.Height)
	xs = make([]int, 0)
	ys = make([]int, 0)
	labels = make([]string, 0)
	totalOctets := int(def.GetTotalPlacementOctets())

	for curr := int(def.GetYAxisOctetsOrigin()); curr < totalOctets; curr += int(def.GetOctetsPerLine()) {
		xs = append(xs, int(offsetX))
		ys = append(ys, int(offsetY+int(dim.Cell.Height)*(curr/int(def.GetOctetsPerLine()))))
		labels = append(labels, fmt.Sprintf("%d", curr))
	}
	return
}

type Cursor struct {
	x uint
	y uint
}

func drawPlacements(def *Definition, dim Dimensions, canvas *svg.SVG) {
	cur := &Cursor{x: 0, y: 0}

	for i, p := range def.Placements {
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
		if p.VariableLength != nil {
			drawBreakMark(def, polygon, canvas)
		}
		drawPlacementText(def, dim, p, polygon, canvas)
	}
}

func drawBreakMark(def *Definition, polygon Polygon, canvas *svg.SVG) {
	left, top, right, bottom := polygon.findBoundingBox()

	sx, sy, cx, cy, px, py, ex, ey := getBreakMarkPoints(def, left, (top+bottom)/2-3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(def, left, (top+bottom)/2+3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)

	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(def, right, (top+bottom)/2-3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
	sx, sy, cx, cy, px, py, ex, ey = getBreakMarkPoints(def, right, (top+bottom)/2+3)
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, `class="breakmark"`)
}

func getBreakMarkPoints(def *Definition, x, y uint) (sx, sy, cx, cy, px, py, ex, ey int) {
	w := def.GetBreakMarkWidth() / 2
	sx = int(x - w)
	cx = int(x)
	px = int(x)
	ex = int(x + w)

	h := def.GetBreakMarkHeight() / 2
	sy = int(y)
	cy = int(y - h)
	py = int(y + h)
	ey = int(y)
	return
}

func drawPlacementText(def *Definition, dim Dimensions, p Placement, polygon Polygon, canvas *svg.SVG) {
	left, top, right, bottom := polygon.findBoundingBox()
	canvas.Text(int(left+right)/2, (int(top+bottom)/2)+(int(dim.Cell.Height)/6), p.Label, `class="placement"`)
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

	if lb == 2 && *p.Bits <= def.GetBitsPerLine() {
		return getSeparatedTwoLinesPlacement(def, dim, cur, p, bitDistributions)
	}

	polygon := getMultipleLinesPlacement(def, dim, cur, p, bitDistributions)
	return []Polygon{polygon}
}

func advanceCursor(def *Definition, bitDistributions []uint, cur *Cursor) {
	bitsPerLine := def.GetBitsPerLine()
	for _, bd := range bitDistributions {
		cur.x += bd
		if cur.x == bitsPerLine {
			cur.x = 0
			cur.y += 1
		}
	}
}

func getBitDistributions(bitsPerLine uint, cur *Cursor, p Placement) []uint {
	bitsToGo := uint(0)
	if p.VariableLength == nil {
		bitsToGo = *p.Bits
	} else {
		bitsToGo = p.VariableLength.MaxBits
	}
	availableInCurrentLine := bitsPerLine - cur.x

	if availableInCurrentLine >= bitsToGo {
		return []uint{bitsToGo}
	}

	result := make([]uint, 0)
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
	return availableBitsInCurrentLine(def, cur) >= *p.Bits
}

func availableBitsInCurrentLine(def *Definition, cur *Cursor) uint {
	return def.GetBitsPerLine() - cur.x
}

func getSingleRectPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement) Polygon {
	left := dim.YAxis.Width + (cur.x * dim.Cell.Width)
	right := left + *p.Bits*dim.Cell.Width
	top := dim.XAxis.Height + (cur.y * dim.Cell.Height)
	bottom := top + dim.Cell.Height
	return createRectPolygon(left, top, right, bottom)
}

func createRectPolygon(left, top, right, bottom uint) Polygon {
	xs := make([]uint, 5)
	xs[0] = left
	xs[1] = right
	xs[2] = right
	xs[3] = left
	xs[4] = left

	ys := make([]uint, 5)
	ys[0] = top
	ys[1] = top
	ys[2] = bottom
	ys[3] = bottom
	ys[4] = top

	return NewPolygon(xs, ys)
}

func doesFitInTwoLines(def *Definition, cur *Cursor, p Placement) bool {
	return *p.Bits-availableBitsInCurrentLine(def, cur) < def.GetBitsPerLine()
}

func getSeparatedTwoLinesPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement, bitDistributions []uint) []Polygon {
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

func getMultipleLinesPlacement(def *Definition, dim Dimensions, cur *Cursor, p Placement, bitDistributions []uint) Polygon {
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
		xs := make([]uint, 7)
		xs[0] = startX
		xs[1] = xs[0] + (bitDistributions[0] * cw)
		xs[2] = xs[1]
		xs[3] = dim.YAxis.Width
		xs[4] = xs[3]
		xs[5] = xs[0]
		xs[6] = xs[0]

		ys := make([]uint, 7)
		ys[0] = startY
		ys[1] = ys[0]
		ys[2] = ys[0] + (ch * uint(len(bitDistributions)))
		ys[3] = ys[2]
		ys[4] = ys[0] + ch
		ys[5] = ys[4]
		ys[6] = ys[0]

		return NewPolygon(xs, ys)
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
	xs := make([]uint, 9)
	xs[0] = startX
	xs[1] = xs[0] + (bitDistributions[0] * cw)
	xs[2] = xs[1]
	xs[3] = dim.YAxis.Width + (bitDistributions[1] * cw)
	xs[4] = xs[3]
	xs[5] = dim.YAxis.Width
	xs[6] = xs[5]
	xs[7] = xs[0]
	xs[8] = xs[0]

	ys := make([]uint, 9)
	ys[0] = startY
	ys[1] = ys[0]
	ys[2] = ys[0] + (uint(len(bitDistributions)-1) * ch)
	ys[3] = ys[2]
	ys[4] = ys[3] + ch
	ys[5] = ys[4]
	ys[6] = ys[0] + ch
	ys[7] = ys[6]
	ys[8] = ys[0]

	return NewPolygon(xs, ys)
}
