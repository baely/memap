package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

type viewer struct {
	// Canvas fields
	canvas     js.Value
	currentMap models.Map

	// Renderer fields
	*canvas.Renderer

	// Panning fields
	holding            bool
	panning            bool
	startLat, startLon float64
	startX, startY     int
}

func newViewer(renderer *canvas.Renderer) *viewer {
	return &viewer{
		Renderer: renderer,
	}
}

func (v *viewer) Init(this js.Value, args []js.Value) interface{} {
	v.canvas = args[0]
	return v
}

func (v *viewer) MouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	v.holding = true

	v.startX = x
	v.startY = y

	v.startLat = v.Lat
	v.startLon = v.Lon

	v.canvas.Get("style").Set("cursor", "grabbing")

	return nil
}

func (v *viewer) MouseMove(this js.Value, args []js.Value) interface{} {
	if !v.holding {
		return nil
	}

	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	deltaX := x - v.startX
	deltaY := y - v.startY

	scaleX, scaleY := util.GetScaleXY(v.Lat, v.Lon, v.Zoom)
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	if !v.panning {
		if deltaX*deltaX+deltaY*deltaY > 25 {
			v.panning = true
		}
	}

	if !v.panning {
		return nil
	}

	v.Lat = v.startLat - deltaLat
	v.Lon = v.startLon - deltaLon

	v.Draw()

	return nil
}

func (v *viewer) MouseUp(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	if !v.panning {
		// Click event
		v.click(x, y)
	}

	v.holding = false
	v.panning = false
	v.canvas.Get("style").Set("cursor", "grab")

	return nil
}

func (v *viewer) MouseLeave(this js.Value, args []js.Value) interface{} {
	v.panning = false
	v.holding = false
	v.canvas.Get("style").Set("cursor", "grab")

	return nil
}

func (v *viewer) Wheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	v.Zoom -= float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if v.Zoom < minZoom {
		v.Zoom = minZoom
	}

	if v.Zoom > maxZoom {
		v.Zoom = maxZoom
	}

	v.Draw()

	return nil
}

func (v *viewer) ButtonPress(this js.Value, args []js.Value) interface{} {
	button := args[0]

	js.Global().Get("console").Call("log", "Button pressed:", button)

	return nil
}

func (v *viewer) click(x, y int) {
	const threshold2 = 144 // 12px squared

	v.SelectedNode = nil
	v.SelectedPath = nil

	// Find the first node
	for _, node := range v.CurrentMap.Nodes {
		nodeX, nodeY := util.TranslateToPosition(v.Lat, v.Lon, v.Zoom, v.Width, v.Height, node.Position)

		dx := x - nodeX
		dy := y - nodeY

		distance2 := dx*dx + dy*dy

		if distance2 <= threshold2 {
			v.SelectedNode = node
			break
		}
	}

	// Find the first path
	for _, path := range v.CurrentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(v.Lat, v.Lon, v.Zoom, v.Width, v.Height, node.Position)
			x2, y2 := util.TranslateToPosition(v.Lat, v.Lon, v.Zoom, v.Width, v.Height, nextNode.Position)

			distance2, _, _ := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 {
				v.SelectedPath = path
				break
			}
		}

		if v.SelectedPath != nil {
			break
		}
	}

	v.Draw()
}

func distance2ToLine(x, y, x1, y1, x2, y2 int) (int, int, int) {
	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 && dy == 0 {
		return (x-x1)*(x-x1) + (y-y1)*(y-y1), x1, y1
	}

	t := float64((x-x1)*dx+(y-y1)*dy) / float64(dx*dx+dy*dy)
	t = max(0, min(1, t))

	nearX := float64(x1) + t*float64(dx)
	nearY := float64(y1) + t*float64(dy)

	fx := float64(x) - nearX
	fy := float64(y) - nearY

	return int(fx*fx + fy*fy), int(nearX), int(nearY)
}
