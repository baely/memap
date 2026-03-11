package interactor

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

type InteractorMenu struct {
	Label    string
	Title    string
	Mode     Mode
	Callback func()
}

type Interactor interface {
	Init() interface{}
	GetMenuItems() []InteractorMenu

	MouseDown(this js.Value, args []js.Value) interface{}
	MouseMove(this js.Value, args []js.Value) interface{}
	MouseUp(this js.Value, args []js.Value) interface{}
	MouseLeave(this js.Value, args []js.Value) interface{}
	Wheel(this js.Value, args []js.Value) interface{}
}

type Mode int

const (
	ModeUnspecified Mode = iota
	ModeViewer
	ModeEdit
	ModeNewNode
	ModeDrawPath
	ModeDemo
)

var registry = map[Mode]Interactor{}

func Init(renderer *canvas.Renderer, canvas js.Value) {
	registry[ModeViewer] = NewViewer(renderer, canvas)
	registry[ModeEdit] = NewEditor(renderer, canvas)
	registry[ModeNewNode] = NewNewNode(renderer, canvas)
	registry[ModeDrawPath] = NewDrawPath(renderer, canvas)
	registry[ModeDemo] = NewDemo(renderer, canvas)
}

func Get(mode Mode) Interactor {
	return registry[mode]
}
