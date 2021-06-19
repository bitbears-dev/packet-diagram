package packetdiagram

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
