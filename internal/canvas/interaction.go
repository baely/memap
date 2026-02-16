package canvas

func (r *Renderer) click(x, y int) {
	const threshold2 = 144 // 12px squared

	r.selectedNode = nil
	r.selectedPath = nil

	// Find the first node
	for _, node := range r.currentMap.Nodes {
		nodeX, nodeY := r.TranslateToPosition(node.Position)

		dx := x - nodeX
		dy := y - nodeY

		distance2 := dx*dx + dy*dy

		if distance2 <= threshold2 {
			r.selectedNode = &node
			break
		}
	}

	// Find the first path
	for _, path := range r.currentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := r.TranslateToPosition(node.Position)
			x2, y2 := r.TranslateToPosition(nextNode.Position)

			distance2 := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 {
				r.selectedPath = &path
				break
			}
		}

		if r.selectedPath != nil {
			break
		}
	}

	r.Draw()
}

func distance2ToLine(x, y, x1, y1, x2, y2 int) int {
	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 && dy == 0 {
		return (x-x1)*(x-x1) + (y-y1)*(y-y1)
	}

	t := float64((x-x1)*dx+(y-y1)*dy) / float64(dx*dx+dy*dy)
	t = max(0, min(1, t))

	nearX := float64(x1) + t*float64(dx)
	nearY := float64(y1) + t*float64(dy)

	fx := float64(x) - nearX
	fy := float64(y) - nearY

	return int(fx*fx + fy*fy)
}
