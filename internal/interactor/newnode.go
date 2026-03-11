package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type newNode struct {
	// Renderer fields
	*canvas.Renderer
}

func NewNewNode(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &newNode{
		Renderer: renderer,
	}
}

func (n *newNode) Init() interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *newNode) GetMenuItems() []InteractorMenu {
	return editorMenuItems(n.CurrentMap)
}

func (n *newNode) MouseDown(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *newNode) MouseMove(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *newNode) MouseUp(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *newNode) MouseLeave(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *newNode) Wheel(this js.Value, args []js.Value) interface{} {
	//TODO implement me
	panic("implement me")
}
