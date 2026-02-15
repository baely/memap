package canvas

func (r *Renderer) clear() {
	//r.ctx.Call("clearRect", 0, 0, r.width, r.height)
	r.batch = append(r.batch, []interface{}{
		"call", "clearRect", 0, 0, r.width, r.height,
	})
}

func (r *Renderer) beginPath() {
	//r.ctx.Call("beginPath")
	r.batch = append(r.batch, []interface{}{
		"call", "beginPath",
	})
}

func (r *Renderer) moveTo(x, y int) {
	//r.ctx.Call("moveTo", x, y)
	r.batch = append(r.batch, []interface{}{
		"call", "moveTo", x, y,
	})
}

func (r *Renderer) lineTo(x, y int) {
	//r.ctx.Call("lineTo", x, y)
	r.batch = append(r.batch, []interface{}{
		"call", "lineTo", x, y,
	})
}

func (r *Renderer) setStrokeStyle(strokeStyle string) {
	//r.ctx.Set("strokeStyle", strokeStyle)
	r.batch = append(r.batch, []interface{}{
		"set", "strokeStyle", strokeStyle,
	})
}

func (r *Renderer) setLineWidth(lineWidth int) {
	//r.ctx.Set("lineWidth", lineWidth)
	r.batch = append(r.batch, []interface{}{
		"set", "lineWidth", lineWidth,
	})
}

func (r *Renderer) setLineCap(lineCap string) {
	//r.ctx.Set("lineCap", lineCap)
	r.batch = append(r.batch, []interface{}{
		"set", "lineCap", lineCap,
	})
}

func (r *Renderer) setLineJoin(lineJoin string) {
	//r.ctx.Set("lineJoin", lineJoin)
	r.batch = append(r.batch, []interface{}{
		"set", "lineJoin", lineJoin,
	})
}

func (r *Renderer) stroke() {
	//r.ctx.Call("stroke")
	r.batch = append(r.batch, []interface{}{
		"call", "stroke",
	})
}
