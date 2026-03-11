package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type drawPath struct {
	// Renderer fields
	*canvas.Renderer
}

func NewDrawPath(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &drawPath{
		Renderer: renderer,
	}
}

func (d *drawPath) Init() interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *drawPath) GetMenuItems() []InteractorMenu {
	return editorMenuItems(d.CurrentMap)
}

func (d *drawPath) MouseDown(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *drawPath) MouseMove(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *drawPath) MouseUp(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *drawPath) MouseLeave(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *drawPath) Wheel(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}
