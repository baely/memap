package canvas

import (
	"math"

	"github.com/baely/memap/internal/models"
)

func (r *Renderer) GetScaleXY() (float64, float64) {
	const a = 1.9743504858348
	const m = 50.0 / 27.0

	scaleX := m * math.Pow(a, r.zoom)
	scaleY := scaleX / math.Cos(r.lat*math.Pi/180.0)

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

func Distance(pos1, pos2 models.Position) float64 {
	const R = 6371e3 // Earth radius in meters

	lat1 := pos1.Latitude * math.Pi / 180.0
	lat2 := pos2.Latitude * math.Pi / 180.0
	deltaLat := (pos2.Latitude - pos1.Latitude) * math.Pi / 180.0
	deltaLon := (pos2.Longitude - pos1.Longitude) * math.Pi / 180.0

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d := R * c

	return d
}

func (r *Renderer) DrawPath(path models.Path) {
	// Draw paths
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		x1, y1 := r.TranslateToPosition(node.Position)
		x2, y2 := r.TranslateToPosition(nextNode.Position)

		r.DrawLine(x1, y1, x2, y2, 24, "white")
	}
}

func (r *Renderer) DrawPathLabel(path models.Path) {
	distance := 0.0
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		// Check if line is visible
		x1, y1 := r.TranslateToPosition(node.Position)
		x2, y2 := r.TranslateToPosition(nextNode.Position)

		if (x1 < 0 && x2 < 0) || (x1 > r.width && x2 > r.width) || (y1 < 0 && y2 < 0) || (y1 > r.height && y2 > r.height) {
			continue
		}

		distance += Distance(node.Position, nextNode.Position)
	}

	midPoint := distance / 2

	currentDistance := 0.0
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		// Check if line is visible
		x1, y1 := r.TranslateToPosition(node.Position)
		x2, y2 := r.TranslateToPosition(nextNode.Position)

		if (x1 < 0 && x2 < 0) || (x1 > r.width && x2 > r.width) || (y1 < 0 && y2 < 0) || (y1 > r.height && y2 > r.height) {
			continue
		}

		currentDistance += Distance(node.Position, nextNode.Position)

		if currentDistance >= midPoint {
			x1, y1 := r.TranslateToPosition(node.Position)
			x2, y2 := r.TranslateToPosition(nextNode.Position)

			midX := (x1 + x2) / 2
			midY := (y1 + y2) / 2

			angle := math.Atan2(float64(y2-y1), float64(x2-x1))

			if angle > math.Pi/2 || angle < -math.Pi/2 {
				angle += math.Pi
			}

			r.DrawText(midX, midY, path.Label, 22, angle, "black")

			break
		}
	}
}

func (r *Renderer) DrawNode(node models.Node) {
	x, y := r.TranslateToPosition(node.Position)
	r.DrawCircle(x, y, 12, "white")
}

func (r *Renderer) DrawNodeLabel(node models.Node) {
	x, y := r.TranslateToPosition(node.Position)

	width := r.MeasureText(node.Label, 22) + 12

	r.DrawRect(x-width/2, y-36, width, 24, "white")
	r.DrawText(x, y-24, node.Label, 22, 0, "black")
}

func (r *Renderer) DrawHighlightedObject() {
	if r.selectedNode != nil {
		node := r.selectedNode
		x, y := r.TranslateToPosition(node.Position)
		r.DrawCircle(x, y, 12, "red")
		return
	}

	if r.selectedPath != nil {
		path := r.selectedPath
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := r.TranslateToPosition(node.Position)
			x2, y2 := r.TranslateToPosition(nextNode.Position)

			r.DrawLine(x1, y1, x2, y2, 24, "red")
		}
	}
}

func (r *Renderer) DrawMap() {
	// Draw terrain

	// Draw roads
	for _, path := range r.currentMap.Paths {
		r.DrawPath(path)
	}

	// Draw nodes
	for _, node := range r.currentMap.Nodes {
		r.DrawNode(node)
	}

	// Draw "highlighted" objects
	r.DrawHighlightedObject()

	// Draw road labels
	for _, path := range r.currentMap.Paths {
		r.DrawPathLabel(path)
	}

	// Draw node labels
	for _, node := range r.currentMap.Nodes {
		r.DrawNodeLabel(node)
	}
}
