package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

type cursorMode int

const (
	cursorModeDefault cursorMode = iota
	cursorModeCreateNode
)

type editor struct {
	// Canvas fields
	canvas js.Value
	//currentMap models.Map

	// Renderer fields
	*canvas.Renderer

	// Panning fields
	holding            bool
	panning            bool
	startLat, startLon float64
	startX, startY     int

	// Edit fields
	cursorX, cursorY int

	// Cursor mode
	cursorMode cursorMode
}

func NewEditor(renderer *canvas.Renderer) Interactor {
	return &editor{
		Renderer: renderer,

		cursorMode: cursorModeDefault,
	}
}

func (e *editor) Init(this js.Value, args []js.Value) interface{} {
	e.canvas = args[0]

	e.canvas.Get("style").Set("cursor", "none")

	return e
}

func (e *editor) GetMenuItems() []InteractorMenu {
	return []InteractorMenu{
		{"⬇️", "download"},
		{"❌", "view"},
	}
}

func (e *editor) snapCursor(event js.Value) (int, int) {
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	if e.panning {
		return x, y
	}

	_, _, newX, newY := e.findClosest(x, y, 652)

	return newX, newY
}

func (e *editor) MouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := e.snapCursor(event)

	e.click(x, y)
	e.DrawCursor(x, y, "red")

	e.holding = true

	e.startX = x
	e.startY = y

	if node, _, _, _ := e.findClosest(x, y, 25); node != e.SelectedNode {
		e.SelectedNode = nil
	}

	if e.SelectedNode != nil {
		e.startLat = e.SelectedNode.Position.Latitude
		e.startLon = e.SelectedNode.Position.Longitude
	} else {
		e.startLat = e.Lat
		e.startLon = e.Lon
	}

	return nil
}

func (e *editor) MouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := e.snapCursor(event)

	e.Draw()
	e.DrawCursor(x, y, "red")

	if !e.holding {
		return nil
	}

	deltaX := x - e.startX
	deltaY := y - e.startY

	scaleX, scaleY := util.GetScaleXY(e.Lat, e.Lon, e.Zoom)
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	if !e.panning {
		if deltaX*deltaX+deltaY*deltaY > 25 {
			e.panning = true
		}
	}

	if !e.panning {
		return nil
	}

	if e.SelectedNode != nil {
		e.moveNode(e.startLat+deltaLat, e.startLon+deltaLon)
	} else {
		e.Lat = e.startLat - deltaLat
		e.Lon = e.startLon - deltaLon
	}

	e.Draw()
	e.DrawCursor(x, y, "red")

	return nil
}

func (e *editor) moveNode(lat, lon float64) {
	e.SelectedNode.Position.Latitude = lat
	e.SelectedNode.Position.Longitude = lon
}

func (e *editor) MouseUp(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x, y := e.snapCursor(event)

	//if !e.panning {
	// Click event
	//e.click(x, y)
	//}

	e.holding = false
	e.panning = false
	//e.canvas.Get("style").Set("cursor", "grab")

	e.DrawCursor(x, y, "red")

	return nil
}

func (e *editor) MouseLeave(this js.Value, args []js.Value) interface{} {
	e.panning = false
	e.holding = false
	//e.canvas.Get("style").Set("cursor", "grab")

	return nil
}

func (e *editor) Wheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	e.Zoom -= float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if e.Zoom < minZoom {
		e.Zoom = minZoom
	}

	if e.Zoom > maxZoom {
		e.Zoom = maxZoom
	}

	e.Draw()

	return nil
}

func (e *editor) ButtonPress(button string) interface{} {
	switch button {
	case "download":
		e.download()
	}

	return nil
}

func (e *editor) download() {
	data, err := models.SerialiseMap(e.CurrentMap, "models", "SampleMap")
	if err != nil {
		js.Global().Get("console").Call("log", err.Error())
	}

	filename := "map.go"

	// Create a Uint8Array from your byte slice
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)

	// Create a blob from the array
	array := js.Global().Get("Array").New(uint8Array)
	blob := js.Global().Get("Blob").New(array)

	// Create an object URL for the blob
	url := js.Global().Get("URL").Call("createObjectURL", blob)

	// Create a temporary anchor element and trigger download
	a := js.Global().Get("document").Call("createElement", "a")
	a.Set("href", url)
	a.Set("download", filename)
	a.Call("click")

	// Clean up the object URL
	js.Global().Get("URL").Call("revokeObjectURL", url)
}

func (e *editor) swapMode(mode int) {
	e.cursorMode = cursorMode(mode)
}

func (e *editor) findClosest(x, y int, threshold2 int) (*models.Node, *models.Path, int, int) {
	minDistance := e.Width*e.Width + e.Height*e.Height

	var closestNode *models.Node
	var closestPath *models.Path
	var closestX, closestY int

	// Find the first node
	for _, node := range e.CurrentMap.Nodes {
		if e.panning && node == e.SelectedNode {
			continue
		}

		nodeX, nodeY := util.TranslateToPosition(e.Lat, e.Lon, e.Zoom, e.Width, e.Height, node.Position)

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

	// Find the first path
	for _, path := range e.CurrentMap.Paths {
		if e.panning && path == e.SelectedPath {
			continue
		}

		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(e.Lat, e.Lon, e.Zoom, e.Width, e.Height, node.Position)
			x2, y2 := util.TranslateToPosition(e.Lat, e.Lon, e.Zoom, e.Width, e.Height, nextNode.Position)

			distance2, nearX, nearY := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 && distance2 <= minDistance {
				//return nil, &path, nearX, nearY
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

func (e *editor) click(x, y int) {
	const threshold2 = 144 // 12px squared

	e.SelectedNode = nil
	e.SelectedPath = nil

	node, path, _, _ := e.findClosest(x, y, threshold2)

	if node != nil {
		e.SelectedNode = node
	} else if path != nil {
		e.SelectedPath = path
	}

	e.Draw()
}
