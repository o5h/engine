package glm

type Vec2I struct{ X, Y int32 }

type Line2I struct {
	Start Vec2I
	End   Vec2I
}

func (l *Line2I) NaiveLineDrawing(plot func(int32, int32)) {
	minX, maxX := MinMaxI(l.End.X, l.Start.X)
	minY, maxY := MinMaxI(l.End.Y, l.Start.Y)
	dx := minX - maxX
	dy := minY - maxY
	if dx == 0 && dy == 0 {
		plot(minX, minY)
		return
	}
	if AbsI(dx) > AbsI(dy) {
		for x := minX; x <= maxX; x++ {
			y := minY + dy*(x-minX)/dx
			plot(x, y)
		}
	} else {
		for y := minY; y <= maxY; y++ {
			x := minX + dx*(y-minY)/dy
			plot(x, y)
		}
	}
}

func Vec2ISum(a, b Vec2I) Vec2I { return Vec2I{X: a.X + b.X, Y: a.Y + b.Y} }
