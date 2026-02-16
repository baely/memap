package canvas

func (r *Renderer) clear() {
	r.batch = append(r.batch, []interface{}{
		"call", "clearRect", 0, 0, r.Width, r.Height,
	})
}

func (r *Renderer) beginPath() {
	r.batch = append(r.batch, []interface{}{
		"call", "beginPath",
	})
}

func (r *Renderer) moveTo(x, y int) {
	r.batch = append(r.batch, []interface{}{
		"call", "moveTo", x, y,
	})
}

func (r *Renderer) lineTo(x, y int) {
	r.batch = append(r.batch, []interface{}{
		"call", "lineTo", x, y,
	})
}

func (r *Renderer) setStrokeStyle(strokeStyle string) {
	r.batch = append(r.batch, []interface{}{
		"set", "strokeStyle", strokeStyle,
	})
}

func (r *Renderer) setLineWidth(lineWidth int) {
	r.batch = append(r.batch, []interface{}{
		"set", "lineWidth", lineWidth,
	})
}

func (r *Renderer) setLineDash(spaces ...float64) {
	r.batch = append(r.batch, []interface{}{
		"call", "setLineDash", spaces,
	})
}

func (r *Renderer) setLineCap(lineCap string) {
	r.batch = append(r.batch, []interface{}{
		"set", "lineCap", lineCap,
	})
}

func (r *Renderer) setLineJoin(lineJoin string) {
	r.batch = append(r.batch, []interface{}{
		"set", "lineJoin", lineJoin,
	})
}

func (r *Renderer) stroke() {
	r.batch = append(r.batch, []interface{}{
		"call", "stroke",
	})
}

func (r *Renderer) arc(x, y, radius int, angleA, angleB float64) {
	r.batch = append(r.batch, []interface{}{
		"call", "arc", x, y, radius, angleA, angleB,
	})
}

func (r *Renderer) setFillStyle(fillStyle string) {
	r.batch = append(r.batch, []interface{}{
		"set", "fillStyle", fillStyle,
	})
}

func (r *Renderer) rect(x, y, width, height int) {
	r.batch = append(r.batch, []interface{}{
		"call", "rect", x, y, width, height,
	})
}

func (r *Renderer) fill() {
	r.batch = append(r.batch, []interface{}{
		"call", "fill",
	})
}

func (r *Renderer) save() {
	r.batch = append(r.batch, []interface{}{
		"call", "save",
	})
}

func (r *Renderer) translate(x, y int) {
	r.batch = append(r.batch, []interface{}{
		"call", "translate", x, y,
	})
}

func (r *Renderer) rotate(angle float64) {
	r.batch = append(r.batch, []interface{}{
		"call", "rotate", angle,
	})
}

func (r *Renderer) setFont(font string) {
	r.batch = append(r.batch, []interface{}{
		"set", "font", font,
	})
}

func (r *Renderer) setTextAlign(textAlign string) {
	r.batch = append(r.batch, []interface{}{
		"set", "textAlign", textAlign,
	})
}

func (r *Renderer) setTextBaseline(textBaseline string) {
	r.batch = append(r.batch, []interface{}{
		"set", "textBaseline", textBaseline,
	})
}

func (r *Renderer) fillText(text string, x, y int) {
	r.batch = append(r.batch, []interface{}{
		"call", "fillText", text, x, y,
	})
}

func (r *Renderer) restore() {
	r.batch = append(r.batch, []interface{}{
		"call", "restore",
	})
}

func (r *Renderer) measureText(text string) float64 {
	return r.ctx.Call("measureText", text).Get("width").Float()
}
