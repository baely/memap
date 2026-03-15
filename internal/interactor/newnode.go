package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

type newNode struct {
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

func NewNewNode(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &newNode{
		canvas:   canvas,
		Renderer: renderer,
	}
}

func (n *newNode) Init() interface{} {
	n.canvas.Get("style").Set("cursor", "crosshair")
	return n
}

func (n *newNode) GetMenuItems() []InteractorMenu {
	return editorMenuItems(n.CurrentMap)
}

func (n *newNode) updateNodeCallback(this js.Value, args []js.Value) interface{} {
	label := args[0].String()
	link := args[1].String()
	description := args[2].String()

	node := n.GetSelectedNode()

	if node == nil {
		return nil
	}

	node.Label = label
	node.Link = link
	node.Description = description

	n.Draw()

	return nil
}

func (n *newNode) MouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	n.holding = true
	n.startX = x
	n.startY = y
	n.startLat = n.Lat
	n.startLon = n.Lon

	return nil
}

func (n *newNode) MouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	n.cursorX = x
	n.cursorY = y

	if !n.holding {
		n.Draw()
		n.DrawCursor(x, y, "green")
		return nil
	}

	deltaX := x - n.startX
	deltaY := y - n.startY

	if !n.panning {
		if deltaX*deltaX+deltaY*deltaY > 25 {
			n.panning = true
		}
	}

	if !n.panning {
		return nil
	}

	scaleX, scaleY := util.GetScaleXY(n.Lat, n.Lon, n.Zoom)
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	n.Lat = n.startLat - deltaLat
	n.Lon = n.startLon - deltaLon

	n.Draw()
	n.DrawCursor(x, y, "green")

	return nil
}

func (n *newNode) MouseUp(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	if !n.panning {
		lat, lon := util.TranslateToLatLon(n.Lat, n.Lon, n.Zoom, n.Width, n.Height, x, y)
		node := &models.Node{
			Label: "",
			Position: &models.Position{
				Latitude:  lat,
				Longitude: lon,
			},
		}
		n.CurrentMap.Nodes = append(n.CurrentMap.Nodes, node)

		// Select the new node and open the edit panel
		n.SetSelectedNode(node, true, n.updateNodeCallback)
	}

	n.holding = false
	n.panning = false

	n.Draw()
	n.DrawCursor(x, y, "green")

	return nil
}

func (n *newNode) MouseLeave(this js.Value, args []js.Value) interface{} {
	n.panning = false
	n.holding = false

	return nil
}

func (n *newNode) Wheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	n.Zoom -= float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if n.Zoom < minZoom {
		n.Zoom = minZoom
	}

	if n.Zoom > maxZoom {
		n.Zoom = maxZoom
	}

	n.Draw()
	n.DrawCursor(n.cursorX, n.cursorY, "green")

	return nil
}
