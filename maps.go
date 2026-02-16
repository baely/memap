package main

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/interactor"
	"github.com/baely/memap/internal/models"
)

var EditMode string

func main() {
	var it interactor.Interactor

	currentMap := models.SampleMap

	editMode := EditMode == "true"

	renderer := canvas.NewRenderer(currentMap)
	it = interactor.New(editMode, renderer)

	js.Global().Set("engine", map[string]interface{}{
		"initRenderer": js.FuncOf(func(this js.Value, args []js.Value) any {
			thisCanvas := args[0]

			renderer.Init(this, args)
			it.Init(this, args)

			thisCanvas.Call("addEventListener", "mousedown", js.FuncOf(it.MouseDown))
			thisCanvas.Call("addEventListener", "mousemove", js.FuncOf(it.MouseMove))
			thisCanvas.Call("addEventListener", "mouseup", js.FuncOf(it.MouseUp))
			thisCanvas.Call("addEventListener", "mouseleave", js.FuncOf(it.MouseLeave))
			thisCanvas.Call("addEventListener", "wheel", js.FuncOf(it.Wheel))

			return nil
		}),
		"updateViewport": js.FuncOf(renderer.UpdateViewport),
		"draw":           js.FuncOf(renderer.DrawFromJS),
		"buttonPress":    js.FuncOf(it.ButtonPress),
	})

	<-make(chan struct{})
}
