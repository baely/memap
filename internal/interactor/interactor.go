package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type Interactor interface {
	Init(this js.Value, args []js.Value) interface{}

	MouseDown(this js.Value, args []js.Value) interface{}
	MouseMove(this js.Value, args []js.Value) interface{}
	MouseUp(this js.Value, args []js.Value) interface{}
	MouseLeave(this js.Value, args []js.Value) interface{}
	Wheel(this js.Value, args []js.Value) interface{}
	ButtonPress(this js.Value, args []js.Value) interface{}
}

func New(editMode bool, renderer *canvas.Renderer) Interactor {
	if editMode {
		return newEditor(renderer)
	}

	return newViewer(renderer)
}
