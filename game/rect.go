package game

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, w, h float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rect) Intersects(otherRect Rect) bool {
	return r.X <= otherRect.MaxX() &&
		otherRect.X <= r.MaxX() &&
		r.Y <= otherRect.MaxY() &&
		otherRect.Y <= r.MaxY()
}
