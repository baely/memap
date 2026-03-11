package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type demo struct {
	// Renderer fields
	*canvas.Renderer
}

func NewDemo(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &demo{
		Renderer: renderer,
	}
}

func (d *demo) Init() interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *demo) GetMenuItems() []InteractorMenu {
	return editorMenuItems(d.CurrentMap)
}

func (d *demo) MouseDown(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *demo) MouseMove(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *demo) MouseUp(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *demo) MouseLeave(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *demo) Wheel(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}
