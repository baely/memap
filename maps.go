package main

import (
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
)

func main() {
	renderer := canvas.NewRenderer()

	js.Global().Set("engine", map[string]interface{}{
		"initRenderer":   js.FuncOf(renderer.Init),
		"updateViewport": js.FuncOf(renderer.UpdateViewport),
		"draw":           js.FuncOf(renderer.DrawFromJS),
	})

	<-make(chan struct{})
}
