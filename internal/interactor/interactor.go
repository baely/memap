package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type InteractorMenu struct {
	Label string
	Title string
}

type Interactor interface {
	Init(this js.Value, args []js.Value) interface{}
	GetMenuItems() []InteractorMenu

	MouseDown(this js.Value, args []js.Value) interface{}
	MouseMove(this js.Value, args []js.Value) interface{}
	MouseUp(this js.Value, args []js.Value) interface{}
	MouseLeave(this js.Value, args []js.Value) interface{}
	Wheel(this js.Value, args []js.Value) interface{}
	ButtonPress(string) interface{}
}

func New(editMode bool, renderer *canvas.Renderer) Interactor {
	if editMode {
		return NewEditor(renderer)
	}

	return NewViewer(renderer)
}
