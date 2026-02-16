package canvas

import (
	"encoding/json"
	"fmt"
	"math"
	"syscall/js"

	"github.com/baely/memap/internal/models"
)

type Renderer struct {
	// Canvas fields
	canvas js.Value
	ctx    js.Value

	// Map fields
	CurrentMap *models.Map

	// Batching fields
	batcher js.Value
	batch   [][]interface{}

	// Viewport fields
	Width, Height int

	// Camera fields
	Lat, Lon float64
	Zoom     float64

	// Interaction fields
	SelectedNode *models.Node
	SelectedPath *models.Path
}

func NewRenderer(m *models.Map) *Renderer {
	return &Renderer{
		batch:      make([][]interface{}, 0),
		CurrentMap: m,

		Lat:  -37.814174,
		Lon:  144.963154,
		Zoom: 16.0,
	}
}

func (r *Renderer) Init(this js.Value, args []js.Value) interface{} {
	r.canvas = args[0]
	r.batcher = args[1]
	r.ctx = args[2]

	return nil
}

func (r *Renderer) UpdateViewport(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}

	r.Width = args[0].Int()
	r.Height = args[1].Int()

	return nil
}

func (r *Renderer) DrawFromJS(this js.Value, args []js.Value) interface{} {
	r.Draw()

	return nil
}

func (r *Renderer) Draw() {
	r.Clear()
	r.DrawMap()
	r.SendBatch()
}

func (r *Renderer) SendBatch() {
	batch, _ := json.Marshal(r.batch)
	r.batcher.Invoke(string(batch))
	r.batch = [][]interface{}{}
}

func (r *Renderer) Clear() {
	r.beginPath()
	r.rect(0, 0, r.Width, r.Height)
	r.setFillStyle("#272727")
	r.fill()
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

func (r *Renderer) DrawText(x, y int, text string, size int, angle float64, style string) {
	r.save()
	r.translate(x, y)
	r.rotate(angle)
	r.setFont(fmt.Sprintf("bold %fpx Arial", float64(size)))
	r.setFillStyle(style)
	r.setTextAlign("center")
	r.setTextBaseline("middle")
	r.fillText(text, 0, 0)
	r.restore()
}

func (r *Renderer) DrawCircle(x, y int, radius int, strokeStyle string) {
	r.beginPath()
	r.arc(x, y, radius, 0, 2*math.Pi)
	r.setFillStyle(strokeStyle)
	r.fill()
}

func (r *Renderer) DrawRect(x, y, width, height int, style string) {
	r.beginPath()
	r.rect(x, y, width, height)
	r.setFillStyle(style)
	r.fill()
}

func (r *Renderer) MeasureText(text string, size int) int {
	font := fmt.Sprintf("bold %fpx Arial", float64(size))
	r.ctx.Set("font", font)

	width := r.measureText(text)

	return int(width)
}
