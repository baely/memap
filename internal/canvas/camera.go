package canvas

import (
	"math"

	"github.com/baely/memap/internal/models"
)

func (r *Renderer) GetScaleXY() (float64, float64) {
	const a = 1.9743504858348
	const m = 50.0 / 27.0

	scaleX := m * math.Pow(a, r.zoom)
	scaleY := scaleX * 1.3

	return scaleX, scaleY
}

func (r *Renderer) TranslateToPosition(pos models.Position) (int, int) {
	scaleX, scaleY := r.GetScaleXY()

	x := int(scaleX*(pos.Longitude-r.lon)) + r.width/2
	y := int(scaleY*(r.lat-pos.Latitude)) + r.height/2
	return x, y
}

func (r *Renderer) TranslateToLatLon(x, y int) (float64, float64) {
	scaleX, scaleY := r.GetScaleXY()

	lon := float64(x-r.width/2)/scaleX + r.lon
	lat := float64(r.height/2-y)/scaleY + r.lat
	return lat, lon
}

func (r *Renderer) DrawPath(path models.Path) {
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		x1, y1 := r.TranslateToPosition(node.Position)
		x2, y2 := r.TranslateToPosition(nextNode.Position)

		r.DrawLine(x1, y1, x2, y2, 10, "black")
	}
}

func (r *Renderer) DrawMap(m models.Map) {
	for _, path := range m.Paths {
		r.DrawPath(path)
	}
}
