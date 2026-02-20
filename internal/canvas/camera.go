package canvas

import (
	"math"

	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

func (r *Renderer) DrawPath(path *models.Path) {
	// Draw paths
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		x1, y1 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
		x2, y2 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, nextNode.Position)

		r.DrawLine(x1, y1, x2, y2, 24, "white")
	}
}

func (r *Renderer) DrawPathLabel(path *models.Path) {
	distance := 0.0
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		// Check if line is visible
		x1, y1 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
		x2, y2 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, nextNode.Position)

		if (x1 < 0 && x2 < 0) || (x1 > r.Width && x2 > r.Width) || (y1 < 0 && y2 < 0) || (y1 > r.Height && y2 > r.Height) {
			continue
		}

		distance += util.Distance(node.Position, nextNode.Position)
	}

	midPoint := distance / 2

	currentDistance := 0.0
	for i, node := range path.Nodes {
		if i == len(path.Nodes)-1 {
			break
		}

		nextNode := path.Nodes[i+1]

		// Check if line is visible
		x1, y1 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
		x2, y2 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, nextNode.Position)

		if (x1 < 0 && x2 < 0) || (x1 > r.Width && x2 > r.Width) || (y1 < 0 && y2 < 0) || (y1 > r.Height && y2 > r.Height) {
			continue
		}

		currentDistance += util.Distance(node.Position, nextNode.Position)

		if currentDistance >= midPoint {
			x1, y1 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
			x2, y2 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, nextNode.Position)

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

func (r *Renderer) DrawNode(node *models.Node) {
	x, y := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
	r.DrawCircle(x, y, 12, "white")
}

func (r *Renderer) DrawNodeLabel(node *models.Node) {
	x, y := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)

	width := r.MeasureText(node.Label, 22) + 12

	r.DrawRect(x-width/2, y-36, width, 24, "white")
	r.DrawText(x, y-24, node.Label, 22, 0, "black")
}

func (r *Renderer) DrawBorder() {
	r.beginPath()
	r.rect(0, 0, r.Width, r.Height)
	r.setStrokeStyle("#D9D9D9")
	r.setLineWidth(50)
	r.stroke()

	// Top left
	r.beginPath()
	r.moveTo(50, 25)
	r.lineTo(25, 25)
	r.lineTo(25, 50)
	r.arc(50, 50, 25, math.Pi, 1.5*math.Pi)
	r.setFillStyle("#D9D9D9")
	r.fill()

	// Top right
	r.beginPath()
	r.moveTo(r.Width-25, 50)
	r.lineTo(r.Width-25, 25)
	r.lineTo(r.Width-50, 25)
	r.arc(r.Width-50, 50, 25, 1.5*math.Pi, 2*math.Pi)
	r.setFillStyle("#D9D9D9")
	r.fill()

	// Bottom right
	r.beginPath()
	r.moveTo(r.Width-50, r.Height-25)
	r.lineTo(r.Width-25, r.Height-25)
	r.lineTo(r.Width-25, r.Height-50)
	r.arc(r.Width-50, r.Height-50, 25, 0, 0.5*math.Pi)
	r.setFillStyle("#D9D9D9")
	r.fill()

	// Bottom left
	r.beginPath()
	r.moveTo(25, r.Height-50)
	r.lineTo(25, r.Height-25)
	r.lineTo(50, r.Height-25)
	r.arc(50, r.Height-50, 25, 0.5*math.Pi, math.Pi)
	r.setFillStyle("#D9D9D9")
	r.fill()
}

func (r *Renderer) DrawHighlightedObject() {
	if r.selectedNode != nil {
		node := r.selectedNode
		x, y := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
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

			x1, y1 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, node.Position)
			x2, y2 := util.TranslateToPosition(r.Lat, r.Lon, r.Zoom, r.Width, r.Height, nextNode.Position)

			r.DrawLine(x1, y1, x2, y2, 24, "red")
		}
	}
}

func (r *Renderer) DrawMap() {
	// Draw terrain

	// Draw roads
	for _, path := range r.CurrentMap.Paths {
		r.DrawPath(path)
	}

	// Draw nodes
	for _, node := range r.CurrentMap.Nodes {
		r.DrawNode(node)
	}

	// Draw "highlighted" objects
	r.DrawHighlightedObject()

	// Draw road labels
	for _, path := range r.CurrentMap.Paths {
		r.DrawPathLabel(path)
	}

	// Draw node labels
	for _, node := range r.CurrentMap.Nodes {
		r.DrawNodeLabel(node)
	}

	r.DrawBorder()
}
