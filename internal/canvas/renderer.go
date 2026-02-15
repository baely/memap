package canvas

import (
	"encoding/json"
	"syscall/js"

	"github.com/baely/memap/internal/models"
)

type Renderer struct {
	// Canvas fields
	canvas js.Value

	// Batching fields
	batcher js.Value
	batch   [][]interface{}

	// Viewport fields
	width, height int

	// Camera fields
	lat, lon float64
	zoom     float64

	// Panning fields
	isPanning          bool
	startLat, startLon float64
	startX, startY     int
}

func NewRenderer() *Renderer {
	return &Renderer{
		batch: make([][]interface{}, 0),

		lat:  -37.814174,
		lon:  144.963154,
		zoom: 15.0,
	}
}

func (r *Renderer) Init(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}

	r.canvas = args[0]
	r.batcher = args[1]

	r.addCanvasListeners()

	return nil
}

func (r *Renderer) UpdateViewport(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}

	r.width = args[0].Int()
	r.height = args[1].Int()

	return nil
}

func (r *Renderer) DrawFromJS(this js.Value, args []js.Value) interface{} {
	r.Draw()

	return nil
}

func (r *Renderer) Draw() {
	r.clear()
	r.DrawMap(models.SampleMap)

	batch, _ := json.Marshal(r.batch)
	r.batcher.Invoke(string(batch))
	r.batch = [][]interface{}{}
}

func (r *Renderer) DrawLine(x0, y0, x1, y1 int, width int, strokeStyle string) {
	r.beginPath()
	r.moveTo(x0, y0)
	r.lineTo(x1, y1)
	r.setStrokeStyle(strokeStyle)
	r.setLineWidth(width)
	r.setLineCap("round")
	r.setLineJoin("round")
	r.stroke()
}
