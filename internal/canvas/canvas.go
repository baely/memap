package canvas

import "syscall/js"

func (r *Renderer) addCanvasListeners() {
	r.canvas.Call("addEventListener", "mousedown", js.FuncOf(r.canvasMouseDown))
	r.canvas.Call("addEventListener", "mousemove", js.FuncOf(r.canvasMouseMove))
	r.canvas.Call("addEventListener", "mouseup", js.FuncOf(r.canvasMouseUp))
	r.canvas.Call("addEventListener", "mouseleave", js.FuncOf(r.canvasMouseLeave))

	r.canvas.Call("addEventListener", "wheel", js.FuncOf(r.canvasWheel))
}

func (r *Renderer) canvasMouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	r.isPanning = true

	r.startX = x
	r.startY = y

	r.startLat = r.lat
	r.startLon = r.lon

	r.canvas.Get("style").Set("cursor", "grabbing")

	return nil
}

func (r *Renderer) canvasMouseMove(this js.Value, args []js.Value) interface{} {
	if !r.isPanning {
		return nil
	}

	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	deltaX := x - r.startX
	deltaY := y - r.startY

	scaleX, scaleY := r.GetScaleXY()
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	r.lat = r.startLat - deltaLat
	r.lon = r.startLon - deltaLon

	r.Draw()

	return nil
}

func (r *Renderer) canvasMouseUp(this js.Value, args []js.Value) interface{} {
	r.isPanning = false
	r.canvas.Get("style").Set("cursor", "grab")

	return nil
}

func (r *Renderer) canvasMouseLeave(this js.Value, args []js.Value) interface{} {
	r.isPanning = false
	r.canvas.Get("style").Set("cursor", "grab")

	return nil
}

func (r *Renderer) canvasWheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	r.zoom += float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if r.zoom < minZoom {
		r.zoom = minZoom
	}

	if r.zoom > maxZoom {
		r.zoom = maxZoom
	}

	r.Draw()

	return nil
}
