package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

type demo struct {
	// Canvas fields
	canvas js.Value

	// Renderer fields
	*canvas.Renderer

	// Panning fields
	holding            bool
	panning            bool
	startLat, startLon float64
	startX, startY     int

	// Cursor fields
	cursorX, cursorY int
}

func NewDemo(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &demo{
		canvas:   canvas,
		Renderer: renderer,
	}
}

func (d *demo) Init() interface{} {
	d.canvas.Get("style").Set("cursor", "none")
	return d
}

func (d *demo) GetMenuItems() []InteractorMenu {
	return editorMenuItems(d.CurrentMap)
}

func (d *demo) snapCursor(event js.Value) (int, int) {
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	if d.panning {
		return x, y
	}

	_, _, newX, newY := d.findClosest(x, y, 625)
	return newX, newY
}

func (d *demo) findClosest(x, y int, threshold2 int) (*models.Node, *models.Path, int, int) {
	minDistance := d.Width*d.Width + d.Height*d.Height

	var closestNode *models.Node
	var closestPath *models.Path
	var closestX, closestY int

	var closestSegmentPath *models.Path
	var closestSegmentIndex int
	_ = closestSegmentPath
	_ = closestSegmentIndex

	// Find the closest node
	for _, node := range d.CurrentMap.Nodes {
		nodeX, nodeY := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)

		dx := x - nodeX
		dy := y - nodeY

		distance2 := dx*dx + dy*dy

		if distance2 <= threshold2 && distance2 < minDistance {
			minDistance = distance2
			closestNode = node
			closestX = nodeX
			closestY = nodeY
		}
	}

	// Find the closest path segment
	for _, path := range d.CurrentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)
			x2, y2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, nextNode.Position)

			distance2, nearX, nearY := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 && distance2 <= minDistance {
				minDistance = distance2
				closestPath = path
				closestX = nearX
				closestY = nearY
				closestNode = nil
			}
		}
	}

	if closestNode == nil && closestPath == nil {
		closestX = x
		closestY = y
	}

	return closestNode, closestPath, closestX, closestY
}

func (d *demo) findClosestSegment(x, y int, threshold2 int) (*models.Path, int) {
	minDistance := d.Width*d.Width + d.Height*d.Height

	var closestPath *models.Path
	closestIndex := -1

	for _, path := range d.CurrentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)
			x2, y2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, nextNode.Position)

			distance2, _, _ := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 && distance2 < minDistance {
				minDistance = distance2
				closestPath = path
				closestIndex = i
			}
		}
	}

	return closestPath, closestIndex
}

func (d *demo) clearSelection() {
	d.SetSelectedNode(nil, false, nil)
	d.SetSelectedPath(nil, false, nil)
}

func (d *demo) deleteNode(node *models.Node) {
	d.clearSelection()
	var newPaths []*models.Path
	for _, path := range d.CurrentMap.Paths {
		var segments [][]*models.Node
		var current []*models.Node

		for _, n := range path.Nodes {
			if n == node {
				if len(current) > 0 {
					segments = append(segments, current)
					current = nil
				}
			} else {
				current = append(current, n)
			}
		}
		if len(current) > 0 {
			segments = append(segments, current)
		}

		for _, seg := range segments {
			if len(seg) >= 2 {
				newPaths = append(newPaths, &models.Path{
					Label: path.Label,
					Type:  path.Type,
					Nodes: seg,
				})
			}
		}
	}
	d.CurrentMap.Paths = newPaths

	for i, n := range d.CurrentMap.Nodes {
		if n == node {
			d.CurrentMap.Nodes = append(d.CurrentMap.Nodes[:i], d.CurrentMap.Nodes[i+1:]...)
			break
		}
	}
}

func (d *demo) deletePathSegment(path *models.Path, segIndex int) {
	d.clearSelection()
	before := path.Nodes[:segIndex+1]
	after := path.Nodes[segIndex+1:]

	for i, p := range d.CurrentMap.Paths {
		if p == path {
			d.CurrentMap.Paths = append(d.CurrentMap.Paths[:i], d.CurrentMap.Paths[i+1:]...)
			break
		}
	}

	if len(before) >= 2 {
		beforeCopy := make([]*models.Node, len(before))
		copy(beforeCopy, before)
		d.CurrentMap.Paths = append(d.CurrentMap.Paths, &models.Path{
			Label: path.Label,
			Type:  path.Type,
			Nodes: beforeCopy,
		})
	}
	if len(after) >= 2 {
		afterCopy := make([]*models.Node, len(after))
		copy(afterCopy, after)
		d.CurrentMap.Paths = append(d.CurrentMap.Paths, &models.Path{
			Label: path.Label,
			Type:  path.Type,
			Nodes: afterCopy,
		})
	}
}

func (d *demo) MouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := d.snapCursor(event)

	d.holding = true
	d.startX = x
	d.startY = y
	d.startLat = d.Lat
	d.startLon = d.Lon

	d.Draw()
	d.DrawCursor(x, y, "red")

	return nil
}

func (d *demo) MouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := d.snapCursor(event)

	d.cursorX = x
	d.cursorY = y

	if !d.holding {
		d.Draw()
		d.DrawCursor(x, y, "red")
		return nil
	}

	deltaX := x - d.startX
	deltaY := y - d.startY

	if !d.panning {
		if deltaX*deltaX+deltaY*deltaY > 25 {
			d.panning = true
		}
	}

	if !d.panning {
		return nil
	}

	scaleX, scaleY := util.GetScaleXY(d.Lat, d.Lon, d.Zoom)
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	d.Lat = d.startLat - deltaLat
	d.Lon = d.startLon - deltaLon

	d.Draw()
	d.DrawCursor(x, y, "red")

	return nil
}

func (d *demo) MouseUp(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := d.snapCursor(event)

	if !d.panning {
		const threshold2 = 144

		node, _, _, _ := d.findClosest(x, y, threshold2)

		if node != nil {
			d.deleteNode(node)
		} else {
			path, segIndex := d.findClosestSegment(x, y, threshold2)
			if path != nil && segIndex >= 0 {
				d.deletePathSegment(path, segIndex)
			}
		}
	}

	d.holding = false
	d.panning = false

	d.Draw()
	d.DrawCursor(x, y, "red")

	return nil
}

func (d *demo) MouseLeave(this js.Value, args []js.Value) interface{} {
	d.panning = false
	d.holding = false

	return nil
}

func (d *demo) Wheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	d.Zoom -= float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if d.Zoom < minZoom {
		d.Zoom = minZoom
	}

	if d.Zoom > maxZoom {
		d.Zoom = maxZoom
	}

	d.Draw()
	d.DrawCursor(d.cursorX, d.cursorY, "red")

	return nil
}
