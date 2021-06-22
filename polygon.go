package packetdiagram

type Polygon struct {
	xs []int
	ys []int
}

func NewPolygon(xs, ys []uint) Polygon {
	xis := make([]int, len(xs))
	yis := make([]int, len(ys))

	for i := range xs {
		xis[i] = int(xs[i])
		yis[i] = int(ys[i])
	}

	return Polygon{
		xs: xis,
		ys: yis,
	}
}

func (p Polygon) findBoundingBox() (left uint, top uint, right uint, bottom uint) {
	left = uint(p.xs[0])
	top = uint(p.ys[0])
	right = uint(p.xs[0])
	bottom = uint(p.ys[0])
	for i := range p.xs {
		if uint(p.xs[i]) < left {
			left = uint(p.xs[i])
		}
		if right < uint(p.xs[i]) {
			right = uint(p.xs[i])
		}
		if uint(p.ys[i]) < top {
			top = uint(p.ys[i])
		}
		if bottom < uint(p.ys[i]) {
			bottom = uint(p.ys[i])
		}
	}
	return
}
