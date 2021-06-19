package packetdiagram

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
			Height: 40,
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
