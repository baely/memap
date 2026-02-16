package canvas

import "math"

func (r *Renderer) DrawCursor(x, y int, style string) {
	const radius = 15

	r.moveTo(x, y)
	r.beginPath()
	r.arc(x, y, radius, 0, 2*math.Pi)
	r.setLineDash(4.89, 13.78, 9.78, 13.78, 9.78, 13.78, 4.89, 0)
	r.setLineWidth(6)
	r.setStrokeStyle(style)
	r.stroke()

	r.SendBatch()
	r.setLineDash(0)
}
