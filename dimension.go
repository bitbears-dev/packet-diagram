package packetdiagram

type Dimensions struct {
	Canvas Dimension
	XAxis  Dimension
	YAxis  Dimension
	Cell   Dimension
}

type Dimension struct {
	Width  uint
	Height uint
}

func calculateDimensions(def *Definition) Dimensions {
	xAxisWidth, xAxisHeight := calculateXAxisDimensions(def)
	yAxisWidth, yAxisHeight := calculateYAxisDimensions(def)
	placementWidth := def.GetCellWidth()*def.GetBitsPerLine() + (def.GetBreakMarkWidth() / 2)
	placementHeight := def.GetCellHeight() * def.GetTotalRows()

	return Dimensions{
		Canvas: Dimension{
			Width:  yAxisWidth + placementWidth,
			Height: xAxisHeight + placementHeight,
		},
		XAxis: Dimension{
			Width:  xAxisWidth,
			Height: xAxisHeight,
		},
		YAxis: Dimension{
			Width:  yAxisWidth,
			Height: yAxisHeight,
		},
		Cell: Dimension{
			Width:  def.GetCellWidth(),
			Height: def.GetCellHeight(),
		},
	}
}

func calculateXAxisDimensions(def *Definition) (w, h uint) {
	w = def.GetCellWidth() * (def.GetBitsPerLine() + uint(2))

	h = 0
	if def.ShouldShowXAxisBits() {
		h += def.GetXAxisBitsHeight()
	}
	if def.ShouldShowXAxisOctets() {
		h += def.GetXAxisOctetsHeight()
	}

	return
}

func calculateYAxisDimensions(def *Definition) (w, h uint) {
	w = 0
	if def.ShouldShowYAxisBits() {
		w += def.GetYAxisBitsWidth()
	}
	if def.ShouldShowYAxisOctets() {
		w += def.GetYAxisOctetsWidth()
	}

	h = def.GetCellHeight() * def.GetTotalRows()

	return
}
