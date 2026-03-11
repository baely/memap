package main

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/interactor"
	"github.com/baely/memap/internal/models"
)

type GeoMap struct {
	ctx js.Value

	*canvas.Renderer
	interactorMode interactor.Mode
}

func (g *GeoMap) updateMode(mode interactor.Mode) interface{} {
	g.interactorMode = mode
	interactor.Get(mode).Init()

	g.Draw()
	g.drawMenu()
	return nil
}

const menuPanelID = "menu-panel"

func (g *GeoMap) drawMenu() {
	menuItems := interactor.Get(g.interactorMode).GetMenuItems()

	document := js.Global().Get("document")

	menuPanel := document.Call("getElementById", menuPanelID)

	for menuPanel.Get("lastChild").Truthy() {
		menuPanel.Call("removeChild", menuPanel.Get("lastChild"))
	}

	for _, menuItem := range menuItems {
		button := document.Call("createElement", "button")
		button.Set("textContent", menuItem.Label)
		button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
			if menuItem.Mode != interactor.ModeUnspecified {
				g.updateMode(menuItem.Mode)
				return nil
			}

			if menuItem.Callback != nil {
				menuItem.Callback()
			}

			return nil
		}))
		menuPanel.Call("appendChild", button)
	}
}

func (g *GeoMap) MouseDown(this js.Value, args []js.Value) interface{} {
	return interactor.Get(g.interactorMode).MouseDown(this, args)
}

func (g *GeoMap) MouseMove(this js.Value, args []js.Value) interface{} {
	return interactor.Get(g.interactorMode).MouseMove(this, args)
}

func (g *GeoMap) MouseUp(this js.Value, args []js.Value) interface{} {
	return interactor.Get(g.interactorMode).MouseUp(this, args)
}

func (g *GeoMap) MouseLeave(this js.Value, args []js.Value) interface{} {
	return interactor.Get(g.interactorMode).MouseLeave(this, args)
}

func (g *GeoMap) Wheel(this js.Value, args []js.Value) interface{} {
	return interactor.Get(g.interactorMode).Wheel(this, args)
}

func main() {
	currentMap := models.SampleMap

	renderer := canvas.NewRenderer(currentMap)

	m := &GeoMap{
		js.ValueOf(nil),
		renderer,
		interactor.ModeViewer,
	}

	js.Global().Set("engine", map[string]interface{}{
		"initRenderer": js.FuncOf(func(this js.Value, args []js.Value) any {
			thisCanvas := args[0]

			interactor.Init(renderer, thisCanvas)

			m.ctx = thisCanvas

			m.Renderer.Init(this, args)

			thisCanvas.Call("addEventListener", "mousedown", js.FuncOf(m.MouseDown))
			thisCanvas.Call("addEventListener", "mousemove", js.FuncOf(m.MouseMove))
			thisCanvas.Call("addEventListener", "mouseup", js.FuncOf(m.MouseUp))
			thisCanvas.Call("addEventListener", "mouseleave", js.FuncOf(m.MouseLeave))
			thisCanvas.Call("addEventListener", "wheel", js.FuncOf(m.Wheel))

			m.drawMenu()

			return nil
		}),
		"updateViewport": js.FuncOf(m.UpdateViewport),
		"draw":           js.FuncOf(m.DrawFromJS),
	})

	<-make(chan struct{})
}
